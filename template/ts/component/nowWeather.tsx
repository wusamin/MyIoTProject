import React from "react";
import Chart from 'chart.js';

import { DrawWeather } from "./weatherIcon"
import { Daily, Hourly, OneCall } from "./model/WeatherForecastResponse"
import { MessageWindow, MessageWindowProps, CastedMessageWindow } from "./character"
import dayjs from "dayjs";

let timerID: number;

const yAxisTempID = "y-axis-temp";
const yAxisHuID = "y-axis-hu";

const CHART_TEXT_COLOR = "rgba(20,20,20,1)";
const CHART_TEXT_SIZE = 15;


const clear = (i: number) => {

}

function fetchWeatherInfo(domain: string): (Promise<OneCall>) {
    const ret = fetch(`${domain}/maid/weatherforecast/owm`)
        .then((v: Response) => {
            const j = v.json()
            console.log(j)
            return j;
        });
    return ret;
}

type ChartDataInfo = {
    dataSet: Chart.ChartDataSets;
    yScale: Chart.ChartYAxe;
    xScale: Chart.ChartXAxe;

    get: (data: (number | number[] | null | undefined)[] | Chart.ChartPoint[] | undefined) => ChartDataInfo;
}

const tempChartDataInfo: ChartDataInfo = {
    dataSet: {
        type: 'line',
        label: '気温(度）',
        data: [35, 34, 37, 35, 34, 35, 34, 25],
        borderColor: "rgba(204,34,34,1)",
        backgroundColor: "rgba(0,0,0,0)",
        yAxisID: yAxisTempID,
    },
    xScale: {
        ticks: {
            fontSize: CHART_TEXT_SIZE,
            fontColor: CHART_TEXT_COLOR,
        }
    },
    yScale: {
        id: yAxisTempID,
        ticks: {
            fontSize: CHART_TEXT_SIZE,
            fontColor: CHART_TEXT_COLOR,
            suggestedMax: 40,
            suggestedMin: 0,
            stepSize: 10,
            callback: (value: any, index: any, values: any) => {
                return `${value}度`
            },
        },
    },
    get: (data: (number | number[] | null | undefined)[] | Chart.ChartPoint[] | undefined) => {
        tempChartDataInfo.dataSet.data = data;
        return tempChartDataInfo;
    }
}

const huChartDataInfo: ChartDataInfo = {
    dataSet: {
        type: 'bar',
        // label: '湿度(度）',

        data: [40, 40, 50, 60, 80, 80, 80, 90],
        backgroundColor: "rgba(65,105,255,0.7)",
        yAxisID: yAxisHuID,

    },
    xScale: {
        ticks: {
            fontSize: CHART_TEXT_SIZE,
            fontColor: CHART_TEXT_COLOR,
        }
    },
    yScale: {
        id: yAxisHuID,
        position: "right",
        ticks: {
            fontSize: CHART_TEXT_SIZE,
            fontColor: CHART_TEXT_COLOR,
            suggestedMax: 100,
            suggestedMin: 0,
            stepSize: 20,
            callback: (value: any, index: any, values: any) => {
                return `${value}%`
            },
        },
        gridLines: {
            drawBorder: false,
            drawOnChartArea: false,
        },
    },
    get: (data: (number | number[] | null | undefined)[] | Chart.ChartPoint[] | undefined) => {
        huChartDataInfo.dataSet.data = data;
        return huChartDataInfo;
    }
}

// const chartData = [tempChartDataInfo, huChartDataInfo];

const chartData = (data: Hourly[]): ChartDataInfo[] => {
    const temp = tempChartDataInfo.get(data.map(p => p.temp));
    const hu = huChartDataInfo.get(data.map(p => p.humidity));
    return [temp, hu];
}

const drawCharts = (chartData: ChartDataInfo[], xLabel: string[]) => {
    const ctx = document.getElementById("moChart") as HTMLCanvasElement;
    ctx.height = 200;

    // Chart.defaults.global.defaultFontColor = "black";
    // Chart.defaults.global.defaultFontSize=13;

    const chartConfig: Chart.ChartConfiguration = {
        type: "bar",
        data: {
            labels: xLabel,
            datasets: [],
        },
        options: {
            title: {
                display: false,
                text: '気温（8月1日~8月7日）',
            },
            scales: {
                yAxes: [],
                xAxes: [{

                }],
            },
            maintainAspectRatio: false,
            responsive: true,
            legend: {
                display: false,
                labels: {
                    fontColor: "black",
                    fontSize: 20,
                }
            },
        }
    }

    chartData.forEach((v, i) => {
        if (!!chartConfig.data?.datasets) {
            chartConfig.data.datasets[i] = v.dataSet;
        }
        if (!!chartConfig.options?.scales?.yAxes) {
            chartConfig.options.scales.yAxes[i] = v.yScale;
        }
    });

    if (!!chartConfig.options?.scales?.xAxes) {
        chartConfig.options.scales.xAxes[0] = chartData[0].xScale;
    }

    new Chart(
        ctx,
        chartConfig
    );
}

const drawChart = () => {
    const ctx = document.getElementById("moChart") as HTMLCanvasElement
    ctx.height = 200;
    new Chart(
        ctx,
        {
            type: 'bar',
            data: {
                labels: ['8月1日', '8月2日', '8月3日', '8月4日', '8月5日', '8月6日', '8月7日'],
                datasets: [
                    {
                        type: 'line',
                        label: '気温(度）',
                        data: [35, 34, 37, 35, 34, 35, 34, 25],
                        borderColor: "rgba(255,0,0,1)",
                        backgroundColor: "rgba(0,0,0,0)",
                        yAxisID: yAxisTempID,
                    },
                    {
                        type: 'bar',
                        label: '湿度(度）',
                        data: [40, 40, 50, 60, 80, 80, 80, 90],
                        backgroundColor: "rgba(0,0,200,0.7)",
                        yAxisID: yAxisHuID,
                    },
                ],
            },
            options: {
                title: {
                    display: true,
                    text: '気温（8月1日~8月7日）'
                },
                scales: {
                    yAxes: [
                        {
                            id: yAxisTempID,
                            ticks: {
                                suggestedMax: 40,
                                suggestedMin: 0,
                                stepSize: 10,
                                callback: (value: any, index: any, values: any) => {
                                    return `${value}度`
                                }
                            }
                        },
                        {
                            id: yAxisHuID,
                            position: "right",
                            ticks: {
                                suggestedMax: 100,
                                suggestedMin: 0,
                                stepSize: 20,
                                callback: (value: any, index: any, values: any) => {
                                    return `${value}%`
                                },
                            },
                            gridLines: {
                                drawBorder: false,
                                drawOnChartArea: false,
                            }
                        },
                    ]
                },
                maintainAspectRatio: false,
                responsive: true,
            }
        }
    )
}

const DrawWeatherIcon: React.FC<OneCall> = (res: OneCall) => {
    if (!res) {
        console.log(`return for res is null.`)
        return <div></div>
    }

    // return (DrawWeather(res.current?.weather[0].main))
    return <div>""</div>
}

const tempArea = (today: Daily) => {
    return (
        <>
            <span className="temp-highest mr-4">
                {`${Math.floor(today.temp.max * 10) / 10}℃`}
            </span>
            <span>{"/"}</span>
            <span className="temp-lowest ml-4">
                {`${Math.floor(today.temp.min * 10) / 10}℃`}
            </span>
        </>
    );
};

const weatherArea = (nowWeather: OneCall) => {
    return (<>
        <div className="w-25 text-center mx-auto">
            {DrawWeatherIcon(nowWeather)}
        </div>
        <div className="font-weight-bold text-center mx-auto mt-n3">
            <span className="weather-text">
                {nowWeather.current?.weather[0].description}
            </span>
        </div>
    </>
    );
}

// const ws = new WebSocket(`ws://192.168.0.55:8080/maid/view/telegram-voice/connect`)

// const castVoice = () => {
//     const [message, setMessage] = React.useState("");
//     // const [ws] = React.useState();

//     ws.onmessage = function (msg) {
//         console.log("casted.")
//         console.log(`message:${msg.data}`)
//         setMessage(msg.data)

//         const audio = document.querySelector("#audio") as HTMLAudioElement
//         audio.src = msg.data
//         audio.play().then(() => {  });
//     }

//     React.useEffect(() => {
//         console.log("message hook was called.")
//         if (message == "") {
//             return
//         }
//         const audio = document.querySelector("#audio") as HTMLAudioElement
//         audio.play().then(() => { setMessage("") });
//     }, [message])
// }

// type WeatherType = "Rainy" | "Sunny" | "Clouds" | "Fog" | "Hurricane" | "Snow" | "Windy";
type WeatherType = String

type NowWeatherProps = {
    domain: string
}

export const NowWeather: React.FC<NowWeatherProps> = (prop: NowWeatherProps) => {
    const [now, setDate] = React.useState(new Date());
    const [weather, setWeather] = React.useState("Clouds" as WeatherType);
    const [nowWeather, setNowWeather] = React.useState({} as OneCall);


    React.useEffect(() => {
        const f = async () => {
            const r = await fetchWeatherInfo(prop.domain);
            setNowWeather(r);
        };
        f();
    }, []);
    // React.useEffect(() => {
    //     console.log("message hook waas called.")
    //     if (message == "") {
    //         return
    //     }
    //     const audio = document.querySelector("#audio") as HTMLAudioElement
    //     audio.play().then(() => { console.log("played.") });
    // }, [message])
    if (!!nowWeather.hourly) {
        drawCharts(chartData(nowWeather.hourly),
            nowWeather.hourly.map(p => {
                // console.log(`map:${dayjs(p.dt * 1000)}`);

                return `${String(dayjs(p.dt * 1000).hour())}時`;
            })
        );
        // const d = dayjs(nowWeather.hourly[0].dt);
        const d = new Date();
        d.setTime(nowWeather.hourly[0].dt)
        // console.log(`${nowWeather.hourly[0].dt}:${new Date(nowWeather.hourly[0].dt * 1000)}`);
    }
    // React.useEffect(() => {
    //     timerID = setInterval(
    //         () => {
    //             setDate(new Date());
    //         }, 1000
    //     );
    //     return () => {
    //         clearInterval(timerID) ;
    //     };
    // }, []);

    
    // castVoice();
    return (
        <div className="back-ground-grad px-0 border border-0">
            {/* <audio id="audio" hidden /> */}
            {/* </audio>
            <button onClick={() => {
                if (message == "") {
                    return
                }
                const audio = document.querySelector("#audio") as HTMLAudioElement
                audio.play().then(() => { setMessage("") });
            }}>test</button> */}
            <div>{now.toString()}</div>
            {weatherArea(nowWeather)}
            <div className="font-weight-bold temp-text row">
                <div className="col-12 text-center">
                    {/* {0 < nowWeather.daily?.length ? tempArea(nowWeather.daily[0]) : ""} */}
                </div>
            </div>
            <div className="font-weight-bold mt-3 row">
                <div className="col-6 text-right d-flex flex-row-reverse align-items-center">
                    <img className="align-middle" width="48" height="48" src="/images/now-weather/umbrella.png"></img>
                </div>
                <div className="col-6 text-left rainy-report-text">
                    <span>
                        {/* {0 < nowWeather.daily?.length ? nowWeather.daily[0].rain : ""} */}
                    </span>
                </div>
            </div>
            <div className="mt-3 w-75 mx-auto weather-chart">
                <canvas id="moChart" className="px-2"></canvas>
            </div>

            {/* <MessageWindow {...{
                firstRow: "日付が変わりました。本日はわたしが、お側で時刻をお知らせしますね。うふふっ",
                secondRow: "",
                thirdRow: "",
            }}
            /> */}
            <CastedMessageWindow />
        </div>
    )
}