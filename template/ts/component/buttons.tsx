import React from "react";

type ButtonProp = {
    value: string;
    text: string;
    stateVal: string;

    clickEvent: (v: number) => void;
}


const button: React.FC<ButtonProp> = (prop: ButtonProp) => {
    const isActive = prop.value === prop.stateVal;
    return (
        <label className={`btn btn-getvalue font-weight-bold btn-outline-secondary focus ${isActive ? 'active' : ''}`}>
            <input
                type="radio"
                name="sensorTerm"
                value={prop.value}
                onFocus={((event: React.ChangeEvent<HTMLInputElement>) => { prop.clickEvent(Number(event.target.value)) })}
                onChange={() => { }}
                checked={isActive}
            />
            {prop.text}
        </label>
    )
}

type ButtonsProp = {
    onClickReload: () => void;
    onClick: (buttonVal: number) => void;
    selectedValue: number;
}

export const Buttons: React.FC<ButtonsProp> = (prop: ButtonsProp) => {
    return (
        <div className="text-center mt-4">
            <input type="button"
                id="button"
                className="btn btn-getvalue btn-outline-secondary font-weight-bold"
                value="Reload"
                onClick={() => prop.onClickReload()}
            />
            <div className="d-inline-block ml-2">
                <div className="btn-group btn-group-toggle"
                    data-toggle="buttons"
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        console.log(`chenged:${event.target.value}`)
                    }}>
                    {button({ value: '1', text: '1 hour', stateVal: String(prop.selectedValue), clickEvent: prop.onClick })}
                    {button({ value: '3', text: '3 hour', stateVal: String(prop.selectedValue), clickEvent: prop.onClick })}
                    {button({ value: '6', text: '6 hour', stateVal: String(prop.selectedValue), clickEvent: prop.onClick })}
                    {button({ value: '12', text: '12 hour', stateVal: String(prop.selectedValue), clickEvent: prop.onClick })}
                    {button({ value: '24', text: '24 hour', stateVal: String(prop.selectedValue), clickEvent: prop.onClick })}
                </div>
            </div>
        </div>
    )
}

