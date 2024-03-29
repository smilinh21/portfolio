import axios from 'axios';

const SendFormData = (data) => axios.post('http://localhost:8000/submit', data);

export default SendFormData