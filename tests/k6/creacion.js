import http from 'k6/http';

export let options = {
    vus: 10000,
    duration: '1s',
};

function generateRandomString(length) {
    let result = '';
    let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let charactersLength = characters.length;
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}

export default function () {
    let uniquePath = generateRandomString(10);  // genera una cadena aleatoria de 10 caracteres
    http.post('https://d1tlewtdyf.execute-api.us-east-1.amazonaws.com/testing/create?url=https://example.com/' + uniquePath);
}
