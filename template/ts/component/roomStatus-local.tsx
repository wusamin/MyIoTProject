import React from "react";
import ReactDOM from "react-dom";

import Chart from 'chart.js';
import dayjs from 'dayjs'
import Comp from "./comp";

import { ChartArea } from "./chartArea"
import { Buttons } from "./buttons"
import { SensorResponse, SensorData } from "./model/SensorData"
import { WeatherForecastResponse ,OneCall} from "./model/WeatherForecastResponse"
import { CastedMessageWindow } from "./character"

type UrlParams = {
    sensorType: string;
    deviceId: string;
}

type ParentState = {
    animate: boolean;

    weatherForecast: OneCall;

    tempChartData: SensorResponse[];
    huChartData: SensorResponse[];
    ilChartData: SensorResponse[];
    co2ChartData: SensorResponse[];
    moChartData: SensorResponse[];

    wsText: string
}

type Charts = {
    teChart?: Chart;
    huChart?: Chart;
    ilChart?: Chart;
    co2Chart?: Chart;
    moChart?: Chart;
}
/**
 * 最新のセンサーの値を取得するURLを生成する
 * @param domain 
 * @param urlParam 
 */
const generateSensorNowUrl = (domain: string, urlParam: UrlParams): string => {
    const selectType = 'daily';

    const now = dayjs();
    const to = now.format('YYYYMMDDHHmmss');
    const from = now.subtract(3, 'm').format('YYYYMMDDHHmmss');

    // return `${domain}/maid/sensorvalue/${urlParam.sensorType}/${selectType}?from=${from}&to=${to}&deviceId=${urlParam.deviceId}`
    return `${domain}/maid/sensor/value/${urlParam.deviceId}/${urlParam.sensorType}?from=${from}&to=${to}`;
}

const generateSensorAPIUrl = (domain: string, urlParam: UrlParams, range: number): string => {
    const selectType = 'daily';

    const now = dayjs();
    const to = now.format('YYYYMMDDHHmmss');
    const from = now.subtract(range, 'h').format('YYYYMMDDHHmmss');

    const deviceId = urlParam && `&deviceId=${urlParam.deviceId}`;
    return `${domain}/maid/sensor/value/${urlParam.deviceId}/${urlParam.sensorType}?from=${from}&to=${to}`;

    // return `${domain}/maid/sensorvalue/${urlParam['sensorType']}/${selectType}?from=${from}&to=${to}${deviceId}`;
}

function fetch2JSON<T = any>(promises: Promise<Response>[]): Promise<T[]> {
    const ret = Promise.all(promises)
        .then(
            async (res) => {
                return await Promise.all(
                    res.map(p => {
                        return p.json();
                    })
                )
            }
        );

    return ret;
}

const shiftSensorArray = (src: SensorData[], add: SensorData): SensorData[] => {
    const newArr = src.slice(1, src.length);
    newArr.push(add)

    return newArr;
}

type RoomStatusProp = {
    domain: string;
}
export class RoomStatus extends React.Component<RoomStatusProp, ParentState> {

    // ws: WebSocket = new WebSocket("ws://localhost:8080/ws");

    intervalId: number = 0;

    dataRange: number = 1;

    constructor(props: RoomStatusProp) {
        super(props);

        this.state =
        {
            animate: false,
            weatherForecast:{},
            tempChartData: [],
            huChartData: [],
            ilChartData: [],
            co2ChartData: [],
            moChartData: [],
            wsText: "",
        }

        // const it = this;

        // this.ws.onmessage = function (this: WebSocket, ev: any): any {
        //     it.setState({
        //         wsText: ev.data,
        //     })

        //     return "";
        // }
    }

    componentDidMount() {
        this.callFetchAll(this.props.domain, this.dataRange);
        this.intervalId = window.setInterval(() => {
            this.callFetchOne(this.props.domain);
        }, 1000 * 60 * 2 + 3000);
    }

    componentWillUnmount() {
        clearInterval(this.intervalId);
    }

    callFetchOne = async (domain: string) => {
        const weatherForecast = await fetch2JSON<OneCall>([fetch(`${domain}/maid/weatherforecast/owm`)]);

        const tempUrls: UrlParams[] = [{ sensorType: 'te', deviceId: 'natureremo' }, { sensorType: 'te', deviceId: 'open-weather-map' }];
        const fetchedTemp = await fetch2JSON<SensorResponse>(tempUrls.map(p => fetch(generateSensorNowUrl(domain, p))));

        this.state.tempChartData[0].data = shiftSensorArray(this.state.tempChartData[0].data, fetchedTemp[0].data[fetchedTemp[0].data.length - 1]);
        this.state.tempChartData[1].data = shiftSensorArray(this.state.tempChartData[1].data, fetchedTemp[1].data[fetchedTemp[1].data.length - 1]);

        const tempBody = this.state.tempChartData;

        const huUrls: UrlParams[] = [{ sensorType: 'hu', deviceId: 'natureremo' }, { sensorType: 'hu', deviceId: 'open-weather-map' }];

        const fetchedHu = await fetch2JSON<SensorResponse>(huUrls.map(p => fetch(generateSensorNowUrl(domain, p))));

        this.state.huChartData[0].data = shiftSensorArray(this.state.huChartData[0].data, fetchedHu[0].data[fetchedHu[0].data.length - 1]);
        this.state.huChartData[1].data = shiftSensorArray(this.state.huChartData[1].data, fetchedHu[1].data[fetchedHu[1].data.length - 1]);

        const huBody = this.state.huChartData;

        const ilUrls = [{ sensorType: 'il', deviceId: 'natureremo' }];
        const fetchedIl = await fetch2JSON<SensorResponse>(ilUrls.map(p => fetch(generateSensorNowUrl(domain, p))));

        this.state.ilChartData[0].data = shiftSensorArray(this.state.ilChartData[0].data, fetchedIl[0].data[fetchedIl[0].data.length - 1]);

        const ilBody = this.state.ilChartData;

        const co2Urls = [{ sensorType: 'co2', deviceId: 'co2mini' }];

        const co2Bodys = await fetch2JSON<SensorResponse>(co2Urls.map(p => fetch(generateSensorNowUrl(domain, p))));
        this.state.co2ChartData[0].data = shiftSensorArray(this.state.co2ChartData[0].data, co2Bodys[0].data[co2Bodys[0].data.length - 1]);
        const co2Body = this.state.co2ChartData;

        const moUrls: UrlParams[] = [{ sensorType: 'mo', deviceId: 'natureremo' }];
        const fetchedMo = await fetch2JSON<SensorResponse>(moUrls.map(p => fetch(generateSensorNowUrl(domain, p))));

        this.state.moChartData[0].data = shiftSensorArray(this.state.moChartData[0].data, fetchedMo[0].data[fetchedMo[0].data.length - 1]);

        const moBody = this.state.moChartData;

        this.setState({
            weatherForecast: weatherForecast[0],
            tempChartData: tempBody,
            huChartData: huBody,
            ilChartData: ilBody,
            co2ChartData: co2Body,
            moChartData: moBody,
        })
    }

    callFetchAll = async (domain: string, timeRange: number) => {
        const weatherForecast = await fetch2JSON<OneCall>([fetch(`${domain}/maid/weatherforecast/owm`)]);

        const tempUrls: UrlParams[] = [{ sensorType: 'te', deviceId: 'natureremo' }, { sensorType: 'te', deviceId: 'open-weather-map' }];
        const tempBody = await fetch2JSON<SensorResponse>(tempUrls.map(p => fetch(generateSensorAPIUrl(domain, p, timeRange))));

        const huUrls: UrlParams[] = [{ sensorType: 'hu', deviceId: 'natureremo' }, { sensorType: 'hu', deviceId: 'open-weather-map' }];
        const huBody = await fetch2JSON<SensorResponse>(huUrls.map(p => fetch(generateSensorAPIUrl(domain, p, timeRange))));

        const ilUrls = [{ sensorType: 'il', deviceId: 'natureremo' }];
        const ilBody = await fetch2JSON<SensorResponse>(ilUrls.map(p => fetch(generateSensorAPIUrl(domain, p, timeRange))));

        const co2Urls = [{ sensorType: 'co2', deviceId: 'co2mini' }];
        const co2Body = await fetch2JSON<SensorResponse>(co2Urls.map(p => fetch(generateSensorAPIUrl(domain, p, timeRange))));

        const moUrls: UrlParams[] = [{ sensorType: 'mo', deviceId: 'natureremo' }];
        const moBody = await fetch2JSON<SensorResponse>(moUrls.map(p => fetch(generateSensorAPIUrl(domain, p, timeRange))));

        this.setState({
            weatherForecast: weatherForecast[0],
            tempChartData: tempBody,
            huChartData: huBody,
            ilChartData: ilBody,
            co2ChartData: co2Body,
            moChartData: moBody,
        })
    }

    render() {
        return (
            <>
                <div className="last-update-text mt-1 ml-2">
                    {`Last Update : ${dayjs().format('YYYY/MM/DD HH:mm:ss')}`}
                </div>
                <Comp
                    animate={this.state.animate}
                    tempChartData={this.state.tempChartData}
                    huChartData={this.state.huChartData}
                    ilChartData={this.state.ilChartData}
                    co2ChartData={this.state.co2ChartData}
                    weatherForecast={this.state.weatherForecast}
                />
                <Buttons
                    onClickReload={() => {
                        this.callFetchAll(this.props.domain, this.dataRange);
                    }

                    }
                    onClick={(v: number) => {
                        this.dataRange = v;
                        this.callFetchAll(this.props.domain, this.dataRange);
                    }}
                    selectedValue={this.dataRange}
                />
                <ChartArea
                    tempChartData={this.state.tempChartData}
                    huChartData={this.state.huChartData}
                    ilChartData={this.state.ilChartData}
                    co2ChartData={this.state.co2ChartData}
                    moChartData={this.state.moChartData}
                />
                <CastedMessageWindow />
            </>
        );
    }
}



