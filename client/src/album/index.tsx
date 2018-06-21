import * as React from 'react';
import { match } from 'react-router-dom';
import { IAlbum } from '../album-thumbnail';
import { ImageCollection } from '../image-collection';
import { GlobalProps } from '..';

interface Props extends GlobalProps {
    match: match<{}>;
}

interface State {
    album: undefined | IAlbum;
    loading: boolean;
}

export class Album extends React.Component<Props, State> {
    /**
     *
     */
    constructor(props: Props) {
        super(props);
        
        this.state = {
            loading: true,
            album: undefined
        }
    }

    componentDidMount() {
        fetch(this.props.apiUrl + this.props.match.url).then(async result => {
            const album = await result.json()
            const state =  album as IAlbum
            this.setState({
                loading: false,
                album: state
            })
        })
        .catch(console.error)
    }

    render() {
        if (!this.state.loading && this.state.album != undefined) {
            const album = this.state.album;
            return(
                <div>
                    <h4>{album.name}</h4>
                    <ImageCollection {...this.props} filenames={album.files} baseUrl={this.props.baseUrl + album.name}/>
                </div>)
        }

        return <div>loading...</div>
    }
}