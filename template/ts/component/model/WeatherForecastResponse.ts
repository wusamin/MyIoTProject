export type WeatherForecastResponse = {
    icon: string;
    nowHumidity: number;
    nowSummary: string;
    nowTemperature: number;
    summary: string;
}

export type OneCall = {
    lat?: number;
    lon?: number;
    timezoneOffset?: number;
    current?: Current;
    hourly?: Hourly[];
    daily?: Daily[];
}

export type Current = {
    dt: number;
    sunrise: number;
    sunset: number;
    temp: number;
    feelsLike: number;
    pressure: number;
    humidity: number;
    dewPoint: number;
    uvi: number;
    clouds: number;
    visibility: number;
    windSpeed: number;
    windDeg: number;
    windGust: number;
    weather: Weather[];
}

export type Hourly = {
    dt: number;
    temp: number;
    feelsLike: number;
    pressure: number;
    humidity: number;
    dewPoint: number;
    uvi: number;
    clouds: number;
    visibility: number;
    windSpeed: number;
    windDeg: number;
    windGust: number;
    weather: Weather[];
    pop: number;
}

export type Daily = {
    dt: number;
    sunrise: number;
    sunset: number;
    temp: DailyTemp;
    feelsLike: DailyFeelsLike;
    pressure: number;
    humidity: number;
    dewPoint: number;
    windSpeed: number;
    windDeg: number;
    weather: Weather[];
    clouds: number;
    pop: number;
    rain: number;
    uvi: number;
}

export type DailyTemp = {
    day: number;
    min: number;
    max: number;
    night: number;
    eve: number;
    morn: number;
}

export type DailyFeelsLike = {
    day: number;
    night: number;
    eve: number;
    morn: number; G
}

export type ForecastOnecallResponse = {
    dt: number;
    temp: number;
    feelsLike: number;
    pressure: number;
    humidity: number;
    dewPoint: number;
    uvi: number;
    clouds: number;
    visibility: number;
    windSpeed: number;
    windDeg: number;
    windGust: number;
    weather: Weather[];
    pop: number;
}

export type Weather = {
    id: number;
    main: string;
    description: string;
    icon: string;
}