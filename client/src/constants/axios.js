import axios from 'axios'
import {BaseURL} from './index'
axios.defaults.baseURL = BaseURL
// axios.defaults.headers.common['Authorization'] = 'Beaver ' + localStorage.getItem('token');
