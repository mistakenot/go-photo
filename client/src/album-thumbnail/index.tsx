import * as React from 'react';

import { Card, CardBody, CardTitle, CardSubtitle, CardText } from 'reactstrap';
import CardImg from 'reactstrap/lib/CardImg';

import './album-thumbnail.css'
import { Link } from 'react-router-dom';

export interface IAlbum {
    name: string
    size: number
    count: number
    url: string
    thumbnail: string
}

export interface IAlbumThumbnailProps {
    album: IAlbum
}

export class AlbumThumbnail extends React.Component<IAlbum, {}> {
    public componentDidMount() {
        
    }

    public render() {
        return (
            <Card className="album-thumbnail">
                <Link to={ "/" + this.props.name }>
                    <CardImg top src={ this.props.thumbnail } alt="Card image cap" />
                    <CardBody>
                        <CardTitle>{ this.props.name }</CardTitle>
                        <CardSubtitle>{ this.props.count } photos</CardSubtitle>
                        <CardText>Some quick example text to build on the card title and make up the bulk of the card's content.</CardText>
                    </CardBody>

                </Link>
            </Card>)
    }
}
