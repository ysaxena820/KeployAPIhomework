document.getElementById('getApi').addEventListener('click', getApi);
const apiKey='3efb2c1f02644803b16160214220108';
function getApi(){
    var CityName = document.getElementById('cityName').value;
    console.log(CityName);
    const url = `http://api.weatherapi.com/v1/current.json?key=${apiKey}&q=${CityName}&aqi=no`;
    console.log(url);
    fetch(url)
    .then((res) => res.json())
    .then((data) => {
        let output = `
        <div>
        <div>
        <p>${data.current.temp_c} deg C</p>
        </div>
        </div>
        `;
        console.log('printing data', data);
        document.getElementById('output').innerHTML = output;
}).catch((err) => console.log(err))
}