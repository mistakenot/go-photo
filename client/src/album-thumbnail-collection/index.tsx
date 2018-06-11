import * as React from 'react';
import { IAlbum, AlbumThumbnail } from '../album-thumbnail';
import { CardDeck } from 'reactstrap';

interface Props {
    albums: IAlbum[]
}

export const AlbumThumbnailCollection = (props: Props) => (
    <CardDeck>
        { props.albums.map(album =>
            <AlbumThumbnail {...album} />)}
    </CardDeck>
)