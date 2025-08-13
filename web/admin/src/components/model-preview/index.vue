<template>
    <AModal :key="info?.id" v-model:visible="visible" width="800px" @open="onModalOpen">
        <ADropdown trigger="hover">
            <AButton type="primary" status="danger">操作控制</AButton>
            <br />
            <template #content>
                <!-- @vue-expect-error -->
                <ADoption v-for="item in controlCommands" :value="item.id">
                    <!-- @vue-expect-error -->
                    {{ item.name }}
                </ADoption>
            </template>
        </ADropdown>
        <div id="three-container" style="height: 600px;"></div>
    </AModal>
</template>

<script lang="ts" setup>
import { ref, nextTick } from "vue";
import * as THREE from "three";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";
import device from "@/api/manage/device";
import { DeviceRow } from "@/api/manage/device/types";
import deviceControl from "@/api/manage/device-control";

const visible = ref(false);
const info = ref<DeviceRow>();

let scene: THREE.Scene,
    camera: THREE.PerspectiveCamera,
    renderer: THREE.WebGLRenderer,
    controls: OrbitControls,
    model: THREE.Group | null = null,
    mixer: THREE.AnimationMixer | null = null;

const clock = new THREE.Clock();

function initThree(container: HTMLElement) {
    scene = new THREE.Scene();
    scene.background = new THREE.Color(0xf0f0f0);

    camera = new THREE.PerspectiveCamera(45, container.clientWidth / container.clientHeight, 0.1, 1000);
    camera.position.set(0, 2, 5);
    camera.lookAt(0, 1, 0);

    renderer = new THREE.WebGLRenderer({ antialias: true });
    renderer.setSize(container.clientWidth, container.clientHeight);
    container.appendChild(renderer.domElement);

    // OrbitControls
    controls = new OrbitControls(camera, renderer.domElement);
    controls.enableDamping = true; // 平滑惯性
    controls.dampingFactor = 0.05;
    controls.enablePan = false; // 禁止平移
    controls.minDistance = 1;
    controls.maxDistance = 20;

    // Lights
    const hemiLight = new THREE.HemisphereLight(0xffffff, 0x444444, 1.5);
    hemiLight.position.set(0, 20, 0);
    scene.add(hemiLight);

    const dirLight = new THREE.DirectionalLight(0xffffff, 2);
    dirLight.position.set(5, 10, 7.5);
    scene.add(dirLight);

    // Grid helper
    const grid = new THREE.GridHelper(10, 20);
    scene.add(grid);

    animate();
}

function animate() {
    requestAnimationFrame(animate);

    const delta = clock.getDelta();
    if (mixer) mixer.update(delta);

    if (model) model.rotation.y += 0.005;

    controls.update();
    renderer.render(scene, camera);
}

function loadModel(url: string) {
    const loader = new GLTFLoader();
    loader.load(
        url,
        (gltf) => {
            if (model) scene.remove(model);
            model = gltf.scene;

            // 自动居中和缩放
            const box = new THREE.Box3().setFromObject(model);
            const size = box.getSize(new THREE.Vector3());
            const center = box.getCenter(new THREE.Vector3());
            model.position.x -= center.x;
            model.position.y -= center.y;
            model.position.z -= center.z;

            // 模型站在地面
            model.position.y += size.y / 2;

            // 缩放模型
            const maxSize = Math.max(size.x, size.y, size.z);
            if (maxSize > 0) model.scale.setScalar(1 / maxSize);

            scene.add(model);

            // 动画
            if (gltf.animations && gltf.animations.length > 0) {
                mixer = new THREE.AnimationMixer(model);
                gltf.animations.forEach((clip) => {
                    const action = mixer!.clipAction(clip);
                    action.play();
                });
            }
        },
        undefined,
        (err) => {
            console.error("模型加载失败", err);
        }
    );
}

function onModalOpen() {
    nextTick(() => {
        const container = document.getElementById("three-container");
        if (!container) return;

        initThree(container);

        const modelUrl = info.value?.modelPath
            ? `/dev${info.value.modelPath}`
            : "/dev/uploads/other/20250813/dc137pfjm5nasnjez1.glb";

        setTimeout(() => {
            loadModel(modelUrl);
        }, 50);
    });
}

const controlCommands = ref([])
function open(deviceId: number) {
    device.read(deviceId).then((res) => {
        info.value = res.data;
        visible.value = true;
    });

    deviceControl.list({
        page: 1,
        pageSize: 100,
        name: "",
        extend: "",
        deviceIds: [deviceId],
        deviceId: 0
    }).then(res => {
        // @ts-expect-error
        controlCommands.value = res.data?.items
    })
}

defineExpose({ open });
</script>

<style scoped>
#three-container {
    width: 100%;
    height: 100%;
    background-color: #f0f0f0;
}
</style>
