import axios from 'axios'
import {BaseURL} from './index'
axios.defaults.baseURL = BaseURL
// axios.defaults.headers.common['Authorization'] = 'Beaver ' + localStorage.getItem('token');
// axios.defaults.headers.common['Content-Type'] = 'application/json';
// axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
