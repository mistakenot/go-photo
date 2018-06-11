function Api() {
    this.url = "localhost:8000";
    this.getOverview = function () {
        return fetch(url).then(async result => {
            return await result.json()
        })
    }
}