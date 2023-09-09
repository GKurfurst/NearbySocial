export async function getUserPosition() {
    try {
        const position = await new Promise((resolve, reject) => {
            navigator.geolocation.getCurrentPosition(resolve, reject);
        });

        const { latitude, longitude } = position.coords;
        return { latitude, longitude };
    } catch (error) {
        throw new Error("无法获取用户位置信息");
    }
}
