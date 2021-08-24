import React from "react";
import Chart from 'chart.js';
import dayjs from 'dayjs'

import { SensorResponse, SensorData } from "./model/SensorData"

type Prop = {
    tempChartData: SensorResponse[];
    ilChartData: SensorResponse[];
    huChartData: SensorResponse[];
    co2ChartData: SensorResponse[];
    moChartData: SensorResponse[];
}

const generateChartData = (bodys: SensorResponse[], chartData: Chart.ChartData): Chart.ChartData => {

    chartData.labels = getXScale(bodys[0]['data']);

    bodys.forEach((p, index) => {
        if (!chartData.datasets) {
            return chartData;
        }
        chartData.datasets[index].data = p['data'].map((p: any) => p.Val)
    })

    return chartData;
}

const getXScale = (apiReturn: SensorData[]): string[] => {
    return apiReturn.map(v => {
        const d = dayjs(v.RecordedAt);
        const month = d.month() + 1;
        const day = d.date();
        const hh = d.hour() < 10 ? `0${d.hour()}` : String(d.hour());
        const mm = d.minute() < 10 ? `0${d.minute()}` : String(d.minute());

        return `${hh}:${mm}`;
    });
};

function assertIsDefined<T>(val: T): asserts val is NonNullable<T> {
    if (val === undefined || val === null) {
        throw new Error(
            `Expected 'val' to be defined, but received ${val}`
        );
    }
}

const drawTempChart = (bodys: SensorResponse[], chart?: Chart): Chart | undefined => {
    try {
        if (!!chart) {
            chart.data = generateChartData(bodys, tempChartSetting.chartData);

            assertIsDefined(chart.options.title);
            assertIsDefined(chart.data.labels)
            chart.options.title.text = `気温（${chart.data.labels[0]}～${chart.data.labels[chart.data.labels.length - 1]}）`
            chart.update();

            return chart;
        } else {
            return new Chart(
                document.getElementById("teChart") as HTMLCanvasElement,
                createChartData(bodys, tempChartSetting)
            );
        }
    } catch (e) {
        return undefined;
    }
};

const drawHuChart = (bodys: SensorResponse[], chart?: Chart): Chart | undefined => {
    try {
        if (!!chart) {
            chart.data = generateChartData(bodys, huChartSetting.chartData);

            assertIsDefined(chart.options.title);
            assertIsDefined(chart.data.labels)
            chart.options.title.text = `湿度（${chart.data.labels[0]}～${chart.data.labels[chart.data.labels.length - 1]}）`

            chart.update();

            return chart;
        } else {
            return new Chart(document.getElementById("huChart") as HTMLCanvasElement,
                createChartData(bodys, huChartSetting));
        }
    } catch (e) {
        return undefined;
    }
}

const drawIlChart = (bodys: SensorResponse[], chart?: Chart): Chart | undefined => {
    try {
        if (!!chart) {
            chart.data = generateChartData(bodys, ilChartSetting.chartData);

            assertIsDefined(chart.options.title);
            assertIsDefined(chart.data.labels)
            chart.options.title.text = `照度（${chart.data.labels[0]}～${chart.data.labels[chart.data.labels.length - 1]}）`

            chart.update();

            return chart;
        } else {
            return new Chart(document.getElementById("ilChart") as HTMLCanvasElement,
                createChartData(bodys, ilChartSetting));
        }
    } catch (e) {
        return undefined;
    }
}

const drawCo2Chart = (bodys: SensorResponse[], chart?: Chart): Chart | undefined => {
    try {
        if (!!chart) {
            chart.data = generateChartData(bodys, co2ChartSetting.chartData);

            assertIsDefined(chart.options.title);
            assertIsDefined(chart.data.labels)
            chart.options.title.text = `CO2濃度（${chart.data.labels[0]}～${chart.data.labels[chart.data.labels.length - 1]}）`

            chart.update();

            return chart;
        } else {
            return new Chart(document.getElementById("co2Chart") as HTMLCanvasElement,
                createChartData(bodys, co2ChartSetting));
        }
    } catch (e) {
        return undefined;
    }
}


const drawMoChart = (bodys: SensorResponse[], chart?: Chart): Chart | undefined => {
    try {
        if (!!chart) {
            chart.data = generateChartData(bodys, moChartSetting.chartData);

            assertIsDefined(chart.options.title);
            assertIsDefined(chart.data.labels)
            chart.options.title.text = `動体検知（${chart.data.labels[0]}～${chart.data.labels[chart.data.labels.length - 1]}）`

            chart.update();

            return chart;
        } else {
            return new Chart(document.getElementById("moChart") as HTMLCanvasElement,
                createChartData(bodys, moChartSetting));
        }
    } catch (e) {
        return undefined;
    }
}

const createChartData = (bodys: SensorResponse[], settings: ChartSettings): Chart.ChartConfiguration => {
    const copiedBody = JSON.parse(JSON.stringify(bodys));
    const copiedSettings = JSON.parse(JSON.stringify(settings));

    const x = getXScale(copiedBody[0]['data']);
    if (!!copiedSettings.options.title) {
        copiedSettings.options.title.text = `${copiedSettings.options.title.text}（${x[0]}～${x[x.length - 1]}）`
    }

    return {
        type: 'line',
        data: generateChartData(copiedBody, copiedSettings.chartData),
        options: copiedSettings.options,
    }
}

type ChartSettings = {
    chartData: Chart.ChartData;
    options: Chart.ChartOptions;
}

const tempChartSetting: ChartSettings = {
    chartData: {
        datasets: [
            {
                label: '室温(度）',
                borderColor: 'rgba(123,104,238,255)',
                backgroundColor: "rgba(0,0,0,0)",
            },
            {
                label: '外気温(度）',
                borderColor: "rgba(220,20,60,255)",
                backgroundColor: 'rgba(0,0,0,0)',
            }
        ],
    },
    options: {
        title: {
            display: true,
            text: '気温',
        },
        scales: {
            yAxes: [{
                ticks: {
                    suggestedMax: 35,
                    suggestedMin: 10,
                    stepSize: 5,
                    callback: (value: any, index: any, values: any) => {
                        return `${value}度`
                    }
                }
            }]
        }
    }
};

const huChartSetting: ChartSettings = {
    chartData: {
        datasets: [
            {
                label: '室内湿度(％）',
                borderColor: 'rgba(0,128,0,1)',
                backgroundColor: "rgba(0,128,0,0)",

            },
            {
                label: '外湿度(％）',
                borderColor: "rgba(0,0,255,255)",
                backgroundColor: "rgba(0,0,0,0)",
            }
        ],
    },
    options: {
        title: {
            display: true,
            // text: `湿度（${x[0]}～${x[x.length - 1]}）`,
            text: '湿度'
        },
        scales: {
            yAxes: [{
                ticks: {
                    suggestedMax: 90,
                    suggestedMin: 20,
                    stepSize: 10,
                    callback: (value: any, index: any, values: any) => {
                        return `${value}％`
                    }
                }
            }]
        }
    }
}

const ilChartSetting: ChartSettings = {
    chartData: {
        datasets: [
            {
                label: '照度(％）',
                borderColor: 'rgba(218,165,32,255)',
                backgroundColor: "rgba(0,0,0,0)"
            },
        ],
    },
    options: {
        title: {
            display: true,
            text: '照度'
        },
        scales: {
            yAxes: [{
                ticks: {
                    suggestedMax: 260,
                    suggestedMin: 0,
                    stepSize: 30,
                    callback: (value, index, values) => {
                        return value
                    }
                }
            }]
        }
    }
}

const co2ChartSetting: ChartSettings = {
    chartData: {
        datasets: [
            {
                label: 'CO2（ppm）',
                borderColor: 'rgba(128,0,0,255)',
                backgroundColor: "rgba(0,0,0,0)"
            },
        ],
    },
    options: {
        title: {
            display: true,
            text: `CO2濃度`
        },
        scales: {
            yAxes: [{
                ticks: {
                    suggestedMax: 1100,
                    suggestedMin: 300,
                    stepSize: 100,
                    callback: (value, index, values) => {
                        return value
                    }
                }
            }]
        }
    }
}

const moChartSetting: ChartSettings = {
    chartData: {
        datasets: [
            {
                label: '動体検知（0-1）',
                borderColor: 'rgba(0,0,139,255)',
                backgroundColor: "rgba(0,0,0,0)"
            },
        ],
    },
    options: {
        title: {
            display: true,
            text: `動体検知`
        },
        scales: {
            yAxes: [{
                ticks: {
                    suggestedMax: 1,
                    suggestedMin: 0,
                    stepSize: 1,
                    callback: (value, index, values) => {
                        return value
                    }
                }
            }]
        }
    }
}

export class ChartArea extends React.Component<Prop> {

    tempChart?: Chart;
    huChart?: Chart;
    ilChart?: Chart;
    co2Chart?: Chart;
    moChart?: Chart;

    constructor(props: Prop) {
        super(props);
    }

    drawChart = () => {
        this.tempChart = drawTempChart(this.props.tempChartData, this.tempChart,);
        this.huChart = drawHuChart(this.props.huChartData, this.huChart,);
        this.ilChart = drawIlChart(this.props.ilChartData, this.ilChart,);
        this.co2Chart = drawCo2Chart(this.props.co2ChartData, this.co2Chart,);
        this.moChart = drawMoChart(this.props.moChartData, this.moChart,);
    }

    render() {
        return (
            <>
                <div className="mt-3 row" >
                    <div className="col-6">
                        <canvas id="teChart" className="px-2"></canvas>
                    </div>
                    <div className="col-6">
                        <canvas id="huChart" className="px-2"></canvas>
                    </div>
                    <div className="col-6">
                        <canvas id="ilChart" className="px-2"></canvas>
                    </div>
                    <div className="col-6">
                        <canvas id="co2Chart" className="px-2"></canvas>
                    </div>
                    <div className="col-6">
                        <canvas id="moChart" className="px-2"></canvas>
                    </div>
                </div>
                {this.drawChart()}
            </>
        );
    }
}