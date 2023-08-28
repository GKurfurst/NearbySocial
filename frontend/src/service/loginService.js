import axios from 'axios'

const checkLogin = async (username, password) => {
    console.log(username, password)
    try {

        axios.defaults.baseURL = 'http://localhost:9090/'

        console.log(username, password)
        const response = await axios.post('/api/login', {
            username : username,
            password : password,
        },{
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        })
        console.log(username, password)
        console.log(response)

    } catch (error) {
        if (error.response.status === 422) {
            alert(error.response.data.message)
        }
    }
}
export default {
    checkLogin
}