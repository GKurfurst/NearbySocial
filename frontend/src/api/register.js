import axios from 'axios'

const register = async (username, telephone, password) => {
    try {
        const response = await axios.post('/api/register', {
            username,
            telephone,
            password
        })
        console.log(response)

    } catch (error) {
        if (error.response.status === 422) {
            alert(error.response.data.message)
        }
    }
}
export default {
    register
}