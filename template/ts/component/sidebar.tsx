import React from "react";
import {
    Link,
} from "react-router-dom";
import ReconnectingWebSocket from "reconnectingwebsocket";


const ws = new ReconnectingWebSocket(`wss://wusamin.jp.ngrok.io/maid/view/telegram-voice/connect`);

// WebSocketへの処理の追加は1回で十分
(() => {
    ws.onmessage = function (msg) {
        const audio = document.querySelector("#audio") as HTMLAudioElement
        audio.src = msg.data
        audio.play()
            .then(
                () => { },
                (reason: any) => { console.log(`reason : ${reason}`) });
    }
})();

type Props = {
    urlRoute: string;
    selectedStatus: string;
    onClick: (selected: string) => void;
}

const sidebar: React.FC<Props> = (props: Props) => {
    return (
        <div className="sidebar sidebar-fixed">
            {/* <button className="sidebar-content text-center"> */}
            {/* <!-- <img className="w-75 h-75" src="/images/weather/clear-day.png"></img> --> */}
            {/* </button> */}
            <div className="sidebar-kiritan text-center">
                <img className="w-100 h-100" src="/images/sidebar/kiritan-color-icon.png"></img>
            </div>
            <Link to={`${props.urlRoute}/roomstatus`}>
                <button
                    className={`sidebar-content text-center ${props.selectedStatus === 'roomstatus' ? 'sidebar-selected' : ''}`}
                    onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => { props.onClick(e.currentTarget.value) }}
                    value='roomstatus'
                >
                    <img className="w-75 h-75" src="/images/sidebar/meter.png"></img>
                </button>
            </Link>
            <Link to={`${props.urlRoute}/weather`}>
                <button
                    className={`sidebar-content text-center ${props.selectedStatus === 'weather' ? 'sidebar-selected' : ''}`}
                    onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
                        props.onClick(e.currentTarget.value)
                    }}
                    value='weather'
                >
                    <img className="w-75 h-75" src="/images/sidebar/weather.png"></img>
                </button>
            </Link>
            <Link to={`${props.urlRoute}/system`}>
                <button
                    className={`sidebar-content text-center ${props.selectedStatus === 'system' ? 'sidebar-selected' : ''}`}
                    onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
                        props.onClick(e.currentTarget.value)
                    }}
                    value='system'
                >
                    <img className="w-75 h-75" src="/images/sidebar/raspberry-pi.png"></img>
                </button>
            </Link>
            <Link to={`${props.urlRoute}/schedule`}>
                <button
                    className={`sidebar-content text-center ${props.selectedStatus === 'schedule' ? 'sidebar-selected' : ''}`}
                    onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
                        props.onClick(e.currentTarget.value)
                    }}
                    value='schedule'
                >
                    <img className="w-75 h-75" src="/images/sidebar/timer.png"></img>
                </button>
            </Link>
            <Link to={`${props.urlRoute}/apiwindow`}>
                <button
                    className={`sidebar-content text-center ${props.selectedStatus === 'apiwindow' ? 'sidebar-selected' : ''}`}
                    onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
                        props.onClick(e.currentTarget.value)
                    }}
                    value='apiwindow'
                >
                    <img className="w-75 h-75" src="/images/sidebar/api-window.png"></img>
                </button>
            </Link>

        </div>
    );
}

export default sidebar;