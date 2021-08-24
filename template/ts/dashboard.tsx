import React, { useState } from "react";
import ReactDOM from "react-dom";
import {
    BrowserRouter,
    Route,
} from "react-router-dom";

import Sidebar from "./component/sidebar";
import { RoomStatus } from "./component/roomStatus-local";
import { NowWeather } from "./component/nowWeather";
import { ApiWindow } from "./component/apiWindow";

const URL_ROUTE = '/maid/view/dashboard';

const DOMAIN = '';

const DrawParent: React.FC = () => {
    const [sidebarStatus, setSidebarStatus] = useState('roomstatus');

    return (
        <div className="row">
            <BrowserRouter>
                <div className="col-1 sidebar pr-0">
                    <Sidebar
                        urlRoute={URL_ROUTE}
                        selectedStatus={sidebarStatus}
                        onClick={(selected: string) => { setSidebarStatus(selected) }}
                    />
                </div>
                <div className="col-11">
                    <Route path={`${URL_ROUTE}`} exact render={() => <RoomStatus domain={DOMAIN} />} />
                    <Route path={`${URL_ROUTE}/roomstatus`} render={() => <RoomStatus domain={DOMAIN} />} />
                    <Route path={`${URL_ROUTE}/weather`} exact render={() => <NowWeather domain={DOMAIN} />} />
                    <Route path={`${URL_ROUTE}/apiwindow`} exact render={() => <ApiWindow domain={DOMAIN} />} />
                </div>
            </BrowserRouter>
        </div>
    )
}

ReactDOM.render(<DrawParent />, document.getElementById('app'));