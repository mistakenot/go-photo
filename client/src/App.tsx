/* tslint:disable */

import * as React from 'react';
import './App.css';
import { Navigation } from './navigation';
import { BrowserRouter, Route } from 'react-router-dom';
import { Home } from './home';
import { Album } from './album';
import { GlobalProps } from '.';

export const App: React.SFC<GlobalProps> = (props) => (
    <BrowserRouter>
        <div className="container">
            <Navigation />
            <hr />
            <br />
            <Route exact path='/' render={(p) => <Home {...props} {...p} />} />
            <Route path='/:album' render={(p) => <Album {...props} {...p} />} />
        </div>
    </BrowserRouter>);

export default App;
