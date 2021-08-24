export class UrlParams {
    public sensorType: string;
    public deviceId: string;
    
    constructor(sensorType: string, deviceId: string) {
        this.sensorType = sensorType;
        this.deviceId = deviceId;
    }
}