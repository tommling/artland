import React, { Component, PropTypes } from 'react'
import { browserHistory, Router } from react-router
import { Provider } from 'react-redux'


class AppContainer extends Component {
	static propTypes = {
		routes : PropTypes.object.isRequired,
		store  : PropTypes.object.isRequired
	}



	render () {
	return (
		<Provider store={store}>
			<div style={{ height: '100%' }}>
				<Router history={browserHistory}
	)

	}
}