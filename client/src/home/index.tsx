import * as React from 'react';
import { match } from 'react-router-dom';
import { IAlbum } from '../album-thumbnail';
import { AlbumThumbnailCollection } from '../album-thumbnail-collection';
import { GlobalProps } from '..';

interface IAlbumOverview {
    albums: IAlbum[]
}


interface Props extends GlobalProps {
    match: match<{}>;
}

export class Home extends React.Component<Props, IAlbumOverview> {
    /**
     *
     */
    constructor(props: Props) {
        super(props);
        
        this.state = { albums: [] };
    }

    public componentDidMount() {
        fetch(this.props.apiUrl).then(async result => {
            const overview = await result.json()
            const state =  overview as IAlbumOverview
            console.log(state)
            this.setState(state)
        })
        .catch(console.error)
    }

    render() {
        return (
            <AlbumThumbnailCollection albums={this.state.albums} />)
    }
}