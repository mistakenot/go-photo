/* tslint:disable */

import * as React from 'react';
import './App.css';
import { Navigation } from './navigation';
import { BrowserRouter, Route } from 'react-router-dom';
import { Home } from './home';
import { Album } from './album';

const URL = "http://localhost:8000/api"

export const App: React.SFC<{}> = () => (
    <BrowserRouter>
        <div className="container">
            <Navigation />
            <hr />
            <br />
            <Route exact path='/' render={(props) => <Home {...props} url={URL}/>} />
            <Route path='/:album' component={Album} />
        </div>
    </BrowserRouter>);

export default App;
