export interface ThresholdRow {
    sensorId?: number;
    alarmLabelId?: number;
    sort?: number;
    template?: string;
    level?: string;
    color?: string;
}


export interface EditAddThreshold {
    sensorId: number;
    thresholds: ThresholdRow[];
}