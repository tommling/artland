import react, { Component } from 'react';
import Button from './Button'; // Import the default button

class DangerButton extends Component {
	render() {
	return <Button color="red" />;
	}
}

export default DangerButton;