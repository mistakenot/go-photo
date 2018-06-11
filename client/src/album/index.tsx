import * as React from 'react';
import { match } from 'react-router-dom';

interface Props {
    match: match<{}>
}

export const Album = (props: Props) => (
    <div>{props.match.url}</div>
)