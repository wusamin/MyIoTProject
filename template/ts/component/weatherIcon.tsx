import React from "react";

const ICON_WIDTH = "180";
const ICON_HEIGHT = "180";

export const DrawWeather = (weather: string) => {
    console.log(`weather type:${weather}`)
    switch (weather) {
        case "Clear":
            return <ClearDay />
        case "Rain":
            return <Rainy />
        case "Clouds":
            return <Cloudy />
        case "Fog":
            return <Fog />
        case "Hurricane":
            return <Hurricane />
        case "Snow":
            return <Snow />
        case "Windy":
            return <Windy />
        case "Mist":
            return <Mist />
        default:
            return <></>
    }
}

export const ClearDay = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <g>
                <path fill="#f59e0b" d="M32 23.36A8.64 8.64 0 1123.36 32 8.66 8.66 0 0132 23.36m0-3A11.64 11.64 0 1043.64 32 11.64 11.64 0 0032 20.36z" />
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M32 15.71V9.5">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M32 48.29v6.21">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M43.52 20.48l4.39-4.39">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M20.48 43.52l-4.39 4.39">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M20.48 20.48l-4.39-4.39">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M43.52 43.52l4.39 4.39">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M15.71 32H9.5">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M48.29 32h6.21">
                    <animate attributeName="strokeDasharray" calcMode="spline" dur="5s" keySplines="0.5 0 0.5 1; 0.5 0 0.5 1" keyTimes="0; .5; 1" repeatCount="indefinite" values="3 6; 6 6; 3 6" />
                </path>
                <animateTransform attributeName="transform" dur="45s" from="0 32 32" repeatCount="indefinite" to="360 32 32" type="rotate" />
            </g>
        </svg>
    )
}

export const ClearNight = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <g>
                <path fill="none" stroke="#72b9d5" strokeLinecap="round" strokeLinejoin="round" strokeWidth="3" d="M46.66 36.2a16.66 16.66 0 01-16.78-16.55 16.29 16.29 0 01.55-4.15A16.56 16.56 0 1048.5 36.1c-.61.06-1.22.1-1.84.1z" />
                <animateTransform attributeName="transform" dur="10s" repeatCount="indefinite" type="rotate" values="-5 32 32;15 32 32;-5 32 32" />
            </g>
        </svg>
    )
}

export const Cloudy = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64" className="cloud">
            <g>
                <path fill="none" strokeLinejoin="round" strokeWidth="3" d="M46.5 31.5h-.32a10.49 10.49 0 00-19.11-8 7 7 0 00-10.57 6 7.21 7.21 0 00.1 1.14A7.5 7.5 0 0018 45.5a4.19 4.19 0 00.5 0v0h28a7 7 0 000-14z" />
                <animateTransform attributeName="transform" dur="7s" repeatCount="indefinite" type="translate" values="-3 0; 3 0; -3 0" />
            </g>
        </svg>
    )
}

export const Fog = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <path fill="none" stroke="#e5e7eb" strokeLinejoin="round" strokeWidth="3" d="M46.5 31.5h-.32a10.49 10.49 0 00-19.11-8 7 7 0 00-10.57 6 7.21 7.21 0 00.1 1.14A7.5 7.5 0 0018 45.5a4.19 4.19 0 00.5 0v0h28a7 7 0 000-14z" />
            <g>
                <path fill="none" stroke="#d1d5db" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M17 58h30" />
                <animateTransform attributeName="transform" begin="0s" dur="5s" repeatCount="indefinite" type="translate" values="-4 0; 4 0; -4 0" />
            </g>
            <g>
                <path fill="none" stroke="#d1d5db" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M17 52h30" />
                <animateTransform attributeName="transform" begin="-4s" dur="5s" repeatCount="indefinite" type="translate" values="-4 0; 4 0; -4 0" />
            </g>
        </svg>
    )
}

export const Hurricane = () => {
    return (
        <div className="mxAuto textCenter">
            <svg className="hurricaneSvg hurricane" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink"
                x="0px" y="0px" viewBox="-437 254.4 85 52.6" xmlSpace="preserve" width={ICON_WIDTH} height={ICON_HEIGHT}>
                <path className="cloud" d="M-361.9,280.5c1.4,0,2.6,0.7,3.4,1.7h1.1c0.4-8.2-5.9-10.8-5.9-10.8c-2.2-1.5-5.4-1-5.4-1
c-0.1-4.1-2.9-7.4-2.9-7.4c-4.7-5.5-10.3-4.9-10.3-4.9c-7.4-0.2-11,5.9-11,5.9c-5.6-4-14.3-2.6-18.2,3.1c-0.7,1.1-1.3,2.2-1.8,3.4
c0,0.1-0.3,1.2-0.4,1.2c3.5-0.6,6.6,1.6,6.6,1.6s1.1-1.1,1.2-1.3c2.4-2.4,5.6-3.6,9-3.2c4.4,0.5,8.5,3,9.9,7.4
c0.1,0.2,0.8,2.4,0.6,2.4c5.3,0.1,7.3,3.6,7.3,3.6h13.4C-364.5,281.2-363.3,280.5-361.9,280.5z"/>
                <path className="cloud" d="M-386,279.6c-0.2,0-0.4,0-0.6,0.1c-0.1-0.8-0.2-1.7-0.4-2.4c-0.3-1-0.8-2-1.4-2.9c-2-2.9-5.3-4.8-9-4.8
c-2.3,0-4.4,0.7-6.1,1.9c-0.6,0.4-1.1,0.8-1.6,1.3c-0.2,0.2-0.5,0.5-0.7,0.8c-0.2,0.3-0.4,0.5-0.6,0.8c-1.8-1.2-3.9-1.9-6.2-1.9
c-5.5,0-10,4-10.8,9.3c-3.5,1-6.1,3.9-6.6,7.6h26.3h12.7h2.3l4.7-6.2c0.6-0.8,1.7-0.9,2.5-0.3s0.9,1.7,0.3,2.5l-3.1,4h0.5h5.6h0.7
c0.1,0,0.2-0.4,0.2-1.1C-377.4,283.5-381.3,279.6-386,279.6z"/>
                <polyline className="lightening" points="-382.8,284.2 -387.9,290.9 -380.6,291.2 -387.9,302 " />
                <path className="line" d="M-426.9,294.4l-5.1,7.3" />
                <path className="line" d="M-420.8,294.4l-5.1,7.3" />
                <path className="line" d="M-415.4,294.4l-5.1,7.3" />
                <path className="line" d="M-409.9,294.4l-5.1,7.3" />
                <path className="line" d="M-404.5,294.4l-5.1,7.3" />
                <path className="line" d="M-399.1,294.4l-5.1,7.3" />
                <path className="line" d="M-393.7,294.4l-5.1,7.3" />
                <path className="line" d="M-388.2,294.4l-5.1,7.3" />
                <g>
                    <path className="littlePath path-1" d="M-374.8,287.2h10.6" />
                    <path className="littlePath path-2" d="M-373.8,289.3h10.9" />
                    <path className="bigPath" d="M-376,288.3c0,0,14,0,14,0c1.7,0,3.1-1.4,3.3-3.1c0-0.5,0-1-0.3-1.4c-0.9-2.3-4.1-2.7-5.6-0.7c-0.4,0.6-0.7,1.3-0.7,1.9" />
                    <path className="littlePath path-3" d="M-364.1,285c0-1.2,1-2.2,2.2-2.2s2.2,1,2.2,2.2c0,1.2-1,2.2-2.2,2.2" />
                </g>
            </svg>
        </div>
    )
}

export const PartlyCloudyDay = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <defs>
                <clipPath id="a">
                    <path fill="none" d="M12 35l-5.28-4.21-2-6 1-7 4-5 5-3h6l5 1 3 3L33 20l-6 4h-6l-3 3v4l-4 2-2 2z" />
                </clipPath>
            </defs>
            <g clipPath="url(#a)">
                <g>
                    <path fill="#f59e0b" d="M19 20.05A3.95 3.95 0 1115.05 24 4 4 0 0119 20.05m0-2A5.95 5.95 0 1025 24a5.95 5.95 0 00-6-5.95z" />
                    <path fill="none" stroke="#f59e0b" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="2" d="M19 15.67V12.5M19 35.5v-3.17M24.89 18.11l2.24-2.24M10.87 32.13l2.24-2.24M13.11 18.11l-2.24-2.24M27.13 32.13l-2.24-2.24M10.67 24H7.5M30.5 24h-3.17" />
                    <animateTransform attributeName="transform" dur="45s" from="0 19.22 24.293" repeatCount="indefinite" to="360 19.22 24.293" type="rotate" />
                </g>
            </g>
            <path fill="none" stroke="#e5e7eb" strokeLinejoin="round" strokeWidth="3" d="M46.5 31.5h-.32a10.49 10.49 0 00-19.11-8 7 7 0 00-10.57 6 7.21 7.21 0 00.1 1.14A7.5 7.5 0 0018 45.5a4.19 4.19 0 00.5 0v0h28a7 7 0 000-14z" />
        </svg>
    )
}

export const PartlyCloudyNight = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <defs>
                <clipPath id="a">
                    <path fill="none" d="M12 35l-5.28-4.21-2-6 1-7 4-5 5-3h6l5 1 3 3L33 20l-6 4h-6l-3 3v4l-4 2-2 2z" />
                </clipPath>
            </defs>
            <g clipPath="url(#a)">
                <g>
                    <path fill="none" stroke="#72b9d5" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M29.33 26.68a10.61 10.61 0 01-10.68-10.54A10.5 10.5 0 0119 13.5a10.54 10.54 0 1011.5 13.11 11.48 11.48 0 01-1.17.07z" />
                    <animateTransform attributeName="transform" dur="10s" repeatCount="indefinite" type="rotate" values="-10 19.22 24.293;10 19.22 24.293;-10 19.22 24.293" />
                </g>
            </g>
            <path fill="none" stroke="#e5e7eb" strokeLinejoin="round" strokeWidth="3" d="M46.5 31.5h-.32a10.49 10.49 0 00-19.11-8 7 7 0 00-10.57 6 7.21 7.21 0 00.1 1.14A7.5 7.5 0 0018 45.5a4.19 4.19 0 00.5 0v0h28a7 7 0 000-14z" />
        </svg>
    )
}

export const Rainy = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <path fill="none" stroke="#e5e7eb" strokeLinecap="round" strokeLinejoin="round" strokeWidth="3" d="M43.67 45.5h2.83a7 7 0 000-14h-.32a10.49 10.49 0 00-19.11-8 7 7 0 00-10.57 6 7.21 7.21 0 00.1 1.14A7.5 7.5 0 0018 45.5a4.19 4.19 0 00.5 0v0" />
            <g>
                <path fill="none" stroke="#2885c7" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="2" d="M24.39 43.03l-.78 4.94" />
                <animateTransform attributeName="transform" dur="0.7s" repeatCount="indefinite" type="translate" values="1 -5; -2 10" />
                <animate attributeName="opacity" dur="0.7s" repeatCount="indefinite" values="0;1;1;0" />
            </g>
            <g>
                <path fill="none" stroke="#2885c7" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="2" d="M31.39 43.03l-.78 4.94" />
                <animateTransform attributeName="transform" begin="-0.4s" dur="0.7s" repeatCount="indefinite" type="translate" values="1 -5; -2 10" />
                <animate attributeName="opacity" begin="-0.4s" dur="0.7s" repeatCount="indefinite" values="0;1;1;0" />
            </g>
            <g>
                <path fill="none" stroke="#2885c7" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="2" d="M38.39 43.03l-.78 4.94" />
                <animateTransform attributeName="transform" begin="-0.2s" dur="0.7s" repeatCount="indefinite" type="translate" values="1 -5; -2 10" />
                <animate attributeName="opacity" begin="-0.2s" dur="0.7s" repeatCount="indefinite" values="0;1;1;0" />
            </g>
        </svg>
    )
}

export const Snow = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <path fill="none" stroke="#e5e7eb" strokeLinecap="round" strokeLinejoin="round" strokeWidth="3" d="M43.67 45.5h2.83a7 7 0 000-14h-.32a10.49 10.49 0 00-19.11-8 7 7 0 00-10.57 6 7.21 7.21 0 00.1 1.14A7.5 7.5 0 0018 45.5a4.19 4.19 0 00.5 0v0" />
            <g>
                <g>
                    <g>
                        <path fill="#72b8d4" d="M24.24 42.67l.24.68a.25.25 0 00.35.14l.65-.31a.26.26 0 01.34.34l-.31.65a.25.25 0 00.14.35l.68.24a.25.25 0 010 .48l-.68.24a.25.25 0 00-.14.35l.31.65a.26.26 0 01-.34.34l-.65-.31a.25.25 0 00-.35.14l-.24.68a.25.25 0 01-.48 0l-.24-.68a.25.25 0 00-.35-.14l-.65.31a.26.26 0 01-.34-.34l.31-.65a.25.25 0 00-.14-.35l-.68-.24a.25.25 0 010-.48l.68-.24a.25.25 0 00.14-.35l-.31-.65a.26.26 0 01.34-.34l.65.31a.25.25 0 00.35-.14l.24-.68a.25.25 0 01.48 0z" />
                        <animateTransform attributeName="transform" dur="9s" repeatCount="indefinite" type="rotate" values="0 24 45; 360 24 45" />
                    </g>
                    <animateTransform attributeName="transform" dur="3s" repeatCount="indefinite" type="translate" values="-3 0; 3 0" />
                </g>
                <animateTransform attributeName="transform" dur="3s" repeatCount="indefinite" type="translate" values="2 -6; -2 12" />
                <animate attributeName="opacity" dur="3s" repeatCount="indefinite" values="0;1;1;1;0" />
            </g>
            <g>
                <g>
                    <g>
                        <path fill="#72b8d4" d="M31.24 42.67l.24.68a.25.25 0 00.35.14l.65-.31a.26.26 0 01.34.34l-.31.65a.25.25 0 00.14.35l.68.24a.25.25 0 010 .48l-.68.24a.25.25 0 00-.14.35l.31.65a.26.26 0 01-.34.34l-.65-.31a.25.25 0 00-.35.14l-.24.68a.25.25 0 01-.48 0l-.24-.68a.25.25 0 00-.35-.14l-.65.31a.26.26 0 01-.34-.34l.31-.65a.25.25 0 00-.14-.35l-.68-.24a.25.25 0 010-.48l.68-.24a.25.25 0 00.14-.35l-.31-.65a.26.26 0 01.34-.34l.65.31a.25.25 0 00.35-.14l.24-.68a.25.25 0 01.48 0z" />
                        <animateTransform attributeName="transform" dur="9s" repeatCount="indefinite" type="rotate" values="0 31 45; 360 31 45" />
                    </g>
                    <animateTransform attributeName="transform" begin="-1s" dur="3s" repeatCount="indefinite" type="translate" values="-3 0; 3 0" />
                </g>
                <animateTransform attributeName="transform" begin="-1s" dur="3s" repeatCount="indefinite" type="translate" values="2 -6; -2 12" />
                <animate attributeName="opacity" begin="-1s" dur="3s" repeatCount="indefinite" values="0;1;1;1;0" />
            </g>
            <g>
                <g>
                    <g>
                        <path fill="#72b8d4" d="M38.24 42.67l.24.68a.25.25 0 00.35.14l.65-.31a.26.26 0 01.34.34l-.31.65a.25.25 0 00.14.35l.68.24a.25.25 0 010 .48l-.68.24a.25.25 0 00-.14.35l.31.65a.26.26 0 01-.34.34l-.65-.31a.25.25 0 00-.35.14l-.24.68a.25.25 0 01-.48 0l-.24-.68a.25.25 0 00-.35-.14l-.65.31a.26.26 0 01-.34-.34l.31-.65a.25.25 0 00-.14-.35l-.68-.24a.25.25 0 010-.48l.68-.24a.25.25 0 00.14-.35l-.31-.65a.26.26 0 01.34-.34l.65.31a.25.25 0 00.35-.14l.24-.68a.25.25 0 01.48 0z" />
                        <animateTransform attributeName="transform" dur="9s" repeatCount="indefinite" type="rotate" values="0 38 45; 360 38 45" />
                    </g>
                    <animateTransform attributeName="transform" begin="-1.5s" dur="3s" repeatCount="indefinite" type="translate" values="-3 0; 3 0" />
                </g>
                <animateTransform attributeName="transform" begin="-1.5s" dur="3s" repeatCount="indefinite" type="translate" values="2 -6; -2 12" />
                <animate attributeName="opacity" begin="-1.5s" dur="3s" repeatCount="indefinite" values="0;1;1;1;0" />
            </g>
        </svg>
    )
}

export const Sleet = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <path fill="none" stroke="#e5e7eb" strokeLinecap="round" strokeLinejoin="round" strokeWidth="3" d="M43.67 45.5h2.83a7 7 0 000-14h-.32a10.49 10.49 0 00-19.11-8 7 7 0 00-10.57 6 7.21 7.21 0 00.1 1.14A7.5 7.5 0 0018 45.5a4.19 4.19 0 00.5 0v0" />
            <g>
                <g>
                    <g>
                        <path fill="#72b8d4" d="M24.24 42.67l.24.68a.25.25 0 00.35.14l.65-.31a.26.26 0 01.34.34l-.31.65a.25.25 0 00.14.35l.68.24a.25.25 0 010 .48l-.68.24a.25.25 0 00-.14.35l.31.65a.26.26 0 01-.34.34l-.65-.31a.25.25 0 00-.35.14l-.24.68a.25.25 0 01-.48 0l-.24-.68a.25.25 0 00-.35-.14l-.65.31a.26.26 0 01-.34-.34l.31-.65a.25.25 0 00-.14-.35l-.68-.24a.25.25 0 010-.48l.68-.24a.25.25 0 00.14-.35l-.31-.65a.26.26 0 01.34-.34l.65.31a.25.25 0 00.35-.14l.24-.68a.25.25 0 01.48 0z" />
                        <animateTransform attributeName="transform" dur="9s" repeatCount="indefinite" type="rotate" values="0 24 45; 360 24 45" />
                    </g>
                    <animateTransform attributeName="transform" dur="3s" repeatCount="indefinite" type="translate" values="-3 0; 3 0" />
                </g>
                <animateTransform attributeName="transform" dur="3s" repeatCount="indefinite" type="translate" values="2 -6; -2 12" />
                <animate attributeName="opacity" dur="3s" repeatCount="indefinite" values="0;1;1;1;0" />
            </g>
            <g>
                <g>
                    <g>
                        <path fill="#72b8d4" d="M38.24 42.67l.24.68a.25.25 0 00.35.14l.65-.31a.26.26 0 01.34.34l-.31.65a.25.25 0 00.14.35l.68.24a.25.25 0 010 .48l-.68.24a.25.25 0 00-.14.35l.31.65a.26.26 0 01-.34.34l-.65-.31a.25.25 0 00-.35.14l-.24.68a.25.25 0 01-.48 0l-.24-.68a.25.25 0 00-.35-.14l-.65.31a.26.26 0 01-.34-.34l.31-.65a.25.25 0 00-.14-.35l-.68-.24a.25.25 0 010-.48l.68-.24a.25.25 0 00.14-.35l-.31-.65a.26.26 0 01.34-.34l.65.31a.25.25 0 00.35-.14l.24-.68a.25.25 0 01.48 0z" />
                        <animateTransform attributeName="transform" dur="9s" repeatCount="indefinite" type="rotate" values="0 38 45; 360 38 45" />
                    </g>
                    <animateTransform attributeName="transform" begin="-1.5s" dur="3s" repeatCount="indefinite" type="translate" values="-3 0; 3 0" />
                </g>
                <animateTransform attributeName="transform" begin="-1.5s" dur="3s" repeatCount="indefinite" type="translate" values="2 -6; -2 12" />
                <animate attributeName="opacity" begin="-1.5s" dur="3s" repeatCount="indefinite" values="0;1;1;1;0" />
            </g>
            <g>
                <path fill="none" stroke="#2885c7" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="2" d="M24.08 45.01l-.16.98" />
                <animateTransform attributeName="transform" dur="1.5s" repeatCount="indefinite" type="translate" values="1 -5; -2 10" />
                <animate attributeName="opacity" dur="1.5s" repeatCount="indefinite" values="0;1;1;0" />
            </g>
            <g>
                <path fill="none" stroke="#2885c7" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="2" d="M31.08 45.01l-.16.98" />
                <animateTransform attributeName="transform" begin="-0.5s" dur="1.5s" repeatCount="indefinite" type="translate" values="1 -5; -2 10" />
                <animate attributeName="opacity" begin="-0.5s" dur="1.5s" repeatCount="indefinite" values="0;1;1;0" />
            </g>
            <g>
                <path fill="none" stroke="#2885c7" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="2" d="M38.08 45.01l-.16.98" />
                <animateTransform attributeName="transform" begin="-1s" dur="1.5s" repeatCount="indefinite" type="translate" values="1 -5; -2 10" />
                <animate attributeName="opacity" begin="-1s" dur="1.5s" repeatCount="indefinite" values="0;1;1;0" />
            </g>
        </svg>
    )
}

export const Windy = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <g>
                <path fill="none" stroke="#e5e7eb" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M43.64 20a5 5 0 113.61 8.46h-35.5M29.14 44a5 5 0 103.61-8.46h-21" />
                <animateTransform attributeName="transform" dur="2s" repeatCount="indefinite" type="translate" values="-8 2; 0 -2; 8 0; 0 1; -8 2" />
            </g>
        </svg>
    )
}

export const Mist = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
            <g>
                <path fill="none" stroke="#e5e7eb" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M17 32h30" />
                <animateTransform attributeName="transform" begin="0s" dur="5s" repeatCount="indefinite" type="translate" values="-4 0; 4 0; -4 0" />
            </g>
            <g>
                <path fill="none" stroke="#e5e7eb" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M17 39h30" />
                <animateTransform attributeName="transform" begin="-2s" dur="5s" repeatCount="indefinite" type="translate" values="-3 0; 3 0; -3 0" />
            </g>
            <g>
                <path fill="none" stroke="#e5e7eb" strokeLinecap="round" strokeMiterlimit="10" strokeWidth="3" d="M17 25h30" />
                <animateTransform attributeName="transform" begin="-4s" dur="5s" repeatCount="indefinite" type="translate" values="-4 0; 4 0; -4 0" />
            </g>
        </svg>
    )
}