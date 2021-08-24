import React from "react";
import ReconnectingWebSocket from "reconnectingwebsocket";

export type MessageWindowProps = {
    firstRow: string;
    secondRow: string;
    thirdRow: string;
}

export const MessageWindow = (prop: MessageWindowProps) => {
    return (
        <>
            <div className="speech-container">
                <div className="speech-bubble speech-text px-2 py-1">
                    <span>{prop.firstRow}</span><br />
                    <span>{prop.secondRow}</span><br />
                    <span>{prop.thirdRow}</span>
                </div>
            </div>
            <img className="kiritan-standing" width="240" height="280" src="/images/now-weather/kiritan-standing.png" />
        </>
    )
}

// hookで持たせるとsetMessageするたびにopenし直すため、外で持つ
const ws = new ReconnectingWebSocket(`wss://wusamin.jp.ngrok.io/maid/view/telegram/connect`)
// const ws = new ReconnectingWebSocket(`ws://localhost:8080/maid/view/telegram/connect`)

export const CastedMessageWindow = () => {
    const [message, setMessage] = React.useState("");

    ws.onmessage = function (msg) {
        setMessage(msg.data)
    }

    ws.onopen = function (ev) {
        this.send("abc")
    }

    return (
        <>
            <div className="speech-container">
                <div className="speech-bubble speech-text px-2 py-1">
                    <span>{message}</span><br />
                </div>
            </div>
            <img className="kiritan-standing" width="240" height="280" src="/images/now-weather/kiritan-standing.png" />
        </>
    )
}