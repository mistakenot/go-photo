import * as React from 'react';
import { Card, CardImg, CardTitle, CardBody } from 'reactstrap';

interface Props {
    url: string;
    filename: string;
}

export class Image extends React.Component<Props> {
    render() {
        return (
            <Card className="album-thumbnail">
                <CardImg top src={ this.props.url } alt="Card image cap" />
                <CardBody>
                    <CardTitle>{ this.props.filename }</CardTitle>
                </CardBody>
            </Card>);
    }
}