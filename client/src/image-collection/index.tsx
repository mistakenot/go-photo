import * as React from 'react';
import { CardDeck } from 'reactstrap';
import { Image } from '../image';
import { GlobalProps } from '..';

interface Props extends GlobalProps {
    filenames: string[];
}

export const ImageCollection = (props: Props) => (
    <CardDeck>
        { props.filenames.map(filename =>
            <Image filename={filename} url={props.baseUrl + '/' + filename} />)}
    </CardDeck>
)