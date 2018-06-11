import * as React from "react";
import {Nav, NavItem} from 'reactstrap';
import { Link } from "react-router-dom";

export const Navigation = () => (
    <Nav>
        <NavItem>
            <Link to="/">
                Home
            </Link>
        </NavItem>
    </Nav>)