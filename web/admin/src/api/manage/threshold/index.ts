import { http } from "@/utils/http";
import { ThresholdRow } from './types';

export default {
    info(sensorId: number) {
        return http<ThresholdRow[]>({
            url: "/manage/threshold/info",
            method: 'post',
            data: {
                id: sensorId
            }
        });
    },
    save(sensorId: number, thresholds: ThresholdRow[]) {
        return http({
            url: '/manage/threshold/save',
            data: {
                sensorId,
                thresholds
            },
            method: 'post'
        })
    }
}