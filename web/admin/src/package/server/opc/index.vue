<script lang='ts' setup>
import tool from '@/utils/tool';
import { computed, reactive, ref } from 'vue';

const visible = ref(false)
const emit = defineEmits(['changeExtend'])

const env = import.meta.env;
const headers = ref({
    "authorization": tool.local.get(env.VITE_APP_TOKEN_PREFIX),
})
const modeOptions = reactive([
    { value: "Invalid", label: "Invalid" },
    { value: "None", label: "None" },
    { value: "Sign", label: "Sign" },
    { value: "SignAndEncrypt", label: "SignAndEncrypt" },
])

const signOptions = reactive([
    // { value: "", label: "为空" },
    { value: "None", label: "None" },
    { value: "Basic128Rsa15", label: "Basic128Rsa15" },
    { value: "Basic256", label: "Basic256" },
    { value: "Basic256Sha256", label: "Basic256Sha256" },
    { value: "Aes128Sha256RsaOaep", label: "Aes128Sha256RsaOaep" },
    { value: "Aes256Sha256RsaPss", label: "Aes256Sha256RsaPss" },
])
const signs = reactive([
    { label: "SecurityPolicyURINone", value: "SecurityPolicyURINone" },
    { label: "SecurityPolicyURIBasic128Rsa15", value: "SecurityPolicyURIBasic128Rsa15" },
    { label: "SecurityPolicyURIBasic256", value: "SecurityPolicyURIBasic256" },
    { label: "SecurityPolicyURIBasic256Sha256", value: "SecurityPolicyURIBasic256Sha256" },
    { label: "SecurityPolicyURI201707", value: "SecurityPolicyURI201707" },
])
const extend = ref({
    policy: "",
    username: "",
    password: "",
    certPath: "",
    keyPath: '',
    mode: ""
})

const defaultCertPath = computed(() => {
    if (extend.value.certPath == "") return
    let f = extend.value.certPath.split("/")

    if (f.length == 0) {
        return []
    }
    return [{
        uid: '-1',
        name: f.length != 0 ? f[f.length - 1] : "",
        url: '/dev' + extend.value.certPath
    }]
})

const defaultKeyPath = computed(() => {
    if (extend.value.keyPath == "") return
    let f = extend.value.keyPath.split("/")
    return [{
        uid: '-1',
        name: f.length != 0 ? f[f.length - 1] : "",
        url: '/dev' + extend.value.keyPath
    }]
})

function successCert(file) {
    if (file.response.code == 0) {
        extend.value.certPath = file.response.data.url
    }
}

function successKey(file) {
    if (file.response.code == 0) {
        extend.value.keyPath = file.response.data.url
    }
}

function handleOk() {
    emit("changeExtend", extend.value)
}

function open(value) {
    visible.value = true
    if (value != undefined && value != null) {
        extend.value = value
    }
}
function onBeforeRemoveCert(fileItem){
    extend.value.certPath = ""
}

function onBeforeRemoveKey(fileItem){
    extend.value.keyPath = ""
}


defineExpose({
    open
})
</script>

<template>
    <AModal @ok="handleOk" v-model:visible="visible">
        <AForm>
            <ARow>
                <ACol>
                    <AFormItem label="模式">
                        <ASelect v-model="extend.mode">
                            <AOption v-for="item in modeOptions" :value="item.value">
                                {{ item.label }}
                            </AOption>
                        </ASelect>
                    </AFormItem>
                </ACol>
            </ARow>
            <ARow>
                <ACol>
                    <AFormItem label="安全策略">
                        <ASelect v-model="extend.policy">
                            <AOption v-for="item in signOptions" :value="item.value">
                                {{ item.label }}
                            </AOption>
                        </ASelect>
                    </AFormItem>
                </ACol>
            </ARow>

            <ARow>
                <ACol>
                    <AFormItem label="用户名">
                        <AInput v-model="extend.username" />
                    </AFormItem>
                </ACol>
            </ARow>

            <ARow>
                <ACol>
                    <AFormItem label="密码">
                        <AInput v-model="extend.password" />
                    </AFormItem>
                </ACol>
            </ARow>

            <ARow>
                <ACol>
                    <AFormItem label="签名">
                        <ASelect>
                            <AOption v-for="item in signs" :value="item.value">
                                {{ item.label }}
                            </AOption>
                        </ASelect>
                    </AFormItem>
                </ACol>
            </ARow>

            <ARow>
                <ACol>
                    <AFormItem label="cert证书">
                        <AUpload :on-before-remove="onBeforeRemoveCert" :key="defaultCertPath" :default-file-list="defaultCertPath" @success="successCert" :headers="headers"
                            with-credentials action="/dev/system/uploadFile" />
                    </AFormItem>
                </ACol>
            </ARow>

            <ARow>
                <ACol>
                    <AFormItem label="key证书">
                        <AUpload :on-before-remove="onBeforeRemoveKey" :key="defaultKeyPath" :default-file-list="defaultKeyPath" @success="successKey" :headers="headers"
                            with-credentials action="/dev/system/uploadFile" />
                    </AFormItem>
                </ACol>
            </ARow>
        </AForm>
    </AModal>
</template>

<style scoped></style>
