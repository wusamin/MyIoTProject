export const batteryAction = async () => {
    const battery = await navigator.getBattery();

    const nowBatteryLevel = battery.level * 100;

    if (3 < nowBatteryLevel) {
        return;
    }

    const url = new URL('https://wusamin.ap.ngrok.io.ap.ngrok.io/maid/dashboard/cast/voice');
    url.searchParams.append('manuscript', `バッテリーの残量が${nowBatteryLevel}％です。画面を消灯します。`);

    await fetch(url.toString()).then();

    fetch('https://wusamin.ap.ngrok.io.ap.ngrok.io/maid/dashboard/screen/off').then();
    
}