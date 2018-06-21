import * as React from 'react';
import * as ReactDOM from 'react-dom';
import App from './App';
import './index.css';
import registerServiceWorker from './registerServiceWorker';

import 'font-awesome/css/font-awesome.min.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'mdbreact/dist/css/mdb.css';

export interface GlobalProps {
    baseUrl: string;
    apiUrl: string;
}

ReactDOM.render(
  <App baseUrl="http://localhost:8000/" apiUrl="http://localhost:8000/api" />,
  document.getElementById('root') as HTMLElement
);
registerServiceWorker();
