import React from "react";

import { SensorResponse } from "./model/SensorData"
import { OneCall } from "./model/WeatherForecastResponse"

type OutdoorsProp = {
    weatherForecast: OneCall;
    animete: boolean;
}

/**
 * Add a scale-up-center animation to DOM specified by className.
 * @param className 
 */
const addAnimation = (className: string) => {
    const target = document.getElementsByClassName(className);

    Array.prototype.forEach.call(target, (element: Element) => {
        element.classList.add('scale-up-center');
        setTimeout(() => {
            element.classList.remove('scale-up-center');
        }, 500);
    });
}

const OutdoorStatus: React.FC<OutdoorsProp> = (prop: OutdoorsProp) => {
    return (
        <>
            <div className="col-5 unittext-frame ml-5">
                <div className="unittext-body">
                    <div className="card-headline border-bottom border-dark d-inline-block">
                        <img className="ml-1 my-1" src="/images/dashboard/park.png" width="28" height="28"></img>
                        <span className="card-headline-text ml-1">Outdoors</span>
                    </div>
                    <div className="row my-2 text-dark font-weight-normal">
                        <div className="col-6">
                            <img className="mt-n1 label-img"
                                src="/images/dashboard/thermometer.png"
                                width="32"
                                height="32"></img>
                            <span
                                id="temperatureText"
                                className="sensorText font-unittext d-inline-block">
                                {`${Number(prop.weatherForecast.current?.temp).toFixed(1)}℃`}
                            </span>
                        </div>
                        <div className="col-6">
                            <img className="mt-n1 label-img"
                                src="/images/dashboard/humidity.png"
                                width="32"
                                height="32"></img>
                            <span id="outHumidityText" className="sensorText d-inline-block ml-2">
                                {`${Number(prop.weatherForecast.current?.humidity).toFixed(1)}％`}
                            </span>
                        </div>
                        <div className="col-12">
                            <img
                                id="weather-icon"
                                className="mt-n2"
                                src={`/images/weather/${convertIconText2URL(prop.weatherForecast.current?.weather?.[0].icon) || 'clear-night'}.png`}
                                width="32"
                                height="32"></img>
                            <span id="weatherText" className="sensorText d-inline-block ml-2">
                                {prop.weatherForecast.current?.weather[0].description}
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}

const convertIconText2URL = (iconText: string | undefined): string => {
    if (!!undefined) return ""

    switch (iconText) {
        case "01d":
            return "clear-day"
        case "01n":
            return "clear-night"
        case "02d":
            return "partly-cloudy-day"
        case "02n":
            return "partly-cloudy-night"
        case "03d":
        case "03n":
        case "04d":
        case "04n":
            return "cloudy"
        case "10d":
        case "10n":
            return "rain"
        case "13d":
        case "13n":
            return "snow"
        default:
            return "clear-day"
    }
}

type HomeStatusProp = {
    tempData: SensorResponse[];
    ilData: SensorResponse[];
    huData: SensorResponse[];
    co2Data: SensorResponse[];
}

const generateTempText = (tempData: SensorResponse[]) => {
    try {
        return `${tempData[0].data[tempData[0].data.length - 1].Val}℃`;
    } catch {
        return 'loading...';
    }
}

const generateHuText = (huData: SensorResponse[]) => {
    try {
        return `${huData[0].data[huData[0].data.length - 1].Val}％`;
    } catch {
        return 'loading...';
    }
}

const generateIlText = (ilData: SensorResponse[]) => {
    try {
        return `${ilData[0].data[ilData[0].data.length - 1].Val}`;
    } catch {
        return 'loading...';
    }
}

const generateCo2Text = (co2Data: SensorResponse[]) => {
    try {
        return `${co2Data[0].data[co2Data[0].data.length - 1].Val}ppm`;
    } catch {
        return 'loading...';
    }
}

const HomeStatus: React.FC<HomeStatusProp> = (prop: HomeStatusProp) => {
    return (
        <div className="col-5 unittext-frame ml-5">
            <div className="unittext-body">
                <div className="card-headline border-bottom border-dark pr-2 d-inline-block">
                    <img className="my-1"
                        src="/images/dashboard/home.png"
                        width="28"
                        height="28"></img>
                    <span className="card-headline card-headline-text ml-1">Home</span>
                </div>
                <div className="row my-2 text-dark font-weight-normal">
                    <div className="col-6">
                        <img className="mt-n1 label-img"
                            src="/images/dashboard/thermometer.png"
                            width="32"
                            height="32"></img>
                        <span id="roomTemperatureText" className="sensorText d-inline-block">
                            {generateTempText(prop.tempData)}
                        </span>
                    </div>
                    <div className="col-6">
                        <img className="mt-n1 label-img"
                            src="/images/dashboard/humidity.png"
                            width="32"
                            height="32"></img>
                        <span id="roomHumidityText" className="sensorText ml-2 d-inline-block">
                            {generateHuText(prop.huData)}
                        </span>
                    </div>
                    <div className="col-6">
                        <img className="mt-n1 label-img"
                            src="/images/dashboard/lightbulb.png"
                            width="32"
                            height="32"></img>
                        <span id="roomIlminateText" className="sensorText d-inline-block">
                            {generateIlText(prop.ilData)}
                        </span>
                    </div>
                    <div className="col-6">
                        <img className="mt-n1 label-img"
                            src="/images/dashboard/co2.png"
                            width="32"
                            height="32"></img>
                        <span id="roomCO2Text" className="sensorText ml-2 d-inline-block">
                            {generateCo2Text(prop.co2Data)}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    );
}

type Prop = {
    tempChartData: SensorResponse[];
    huChartData: SensorResponse[];
    ilChartData: SensorResponse[];
    co2ChartData: SensorResponse[];

    weatherForecast: OneCall;

    animate: boolean;
}

const Comp: React.FC<Prop> = (props: Prop) => {
    return (
        <>
            <div className="row mt-4">
                <OutdoorStatus
                    animete={props.animate}
                    weatherForecast={props.weatherForecast}
                />
                <HomeStatus
                    tempData={props.tempChartData}
                    huData={props.huChartData}
                    ilData={props.ilChartData}
                    co2Data={props.co2ChartData}
                />
            </div>
            {addAnimation('sensorText')}
        </>
    );
}

export default Comp;