export type SensorResponse = {
    data: SensorData[];
}

export type SensorData = {
    RecordedAt: Date;
    DeviceID: string;
    SensorType: string;
    Val: string;
    CreatedAt: Date;
}

