import './user-label.css';

import React from 'react';
import PropTypes from 'prop-types';
import ButtonGroup from 'react-bootstrap/ButtonGroup';
import Button from 'react-bootstrap/Button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import UserConfig from '../core/user.js';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import api from '../core/api-service.js';
import axios from 'axios';
import VeloForm from '../forms/form.js';

class UserSettings extends React.PureComponent {
    static contextType = UserConfig;
    static propTypes = {
        onClose: PropTypes.func.isRequired,
        setSetting: PropTypes.func.isRequired,
    }

    componentDidMount = () => {
        if (this.context.traits) {
            this.setState({
                theme: this.context.traits.theme || "no-theme",
                default_password: this.context.traits.default_password || "",
            });
        }
    }

    state = {
        theme: "",
        default_password: "",
    }

    saveSettings = ()=> {
        this.props.setSetting({
            theme: this.state.theme,
            default_password: this.state.default_password,
        });
    }

    render() {
        return (
            <Modal show={true}
                   dialogClassName="modal-70w"
                   onHide={this.props.onClose}>
              <Modal.Header closeButton>
                <Modal.Title>User Settings</Modal.Title>
              </Modal.Header>
              <Modal.Body>
                <Form.Group as={Row}>
                  <Form.Label column sm="3">
                    Theme
                  </Form.Label>
                  <Col sm="8">
                    <Form.Control as="select"
                                  value={this.state.theme}
                                  placeholder="Select a theme"
                                  onChange={(e) => {
                                      this.setState({theme: e.currentTarget.value});

                                      // Change the theme instantly
                                      this.props.setSetting({
                                          theme: e.currentTarget.value,
                                          default_password: this.state.default_password,
                                      });
                                  }}>
                      <option value="no-theme">Default Velociraptor</option>
                      <option value="veloci-light">Velociraptor (light) [experimental]</option>
                      <option value="veloci-dark">Velociraptor (dark) [experimental]</option>
                      {/* <option value="github-dimmed-light">Github dimmed (light)</option> */}
                      <option value="github-dimmed-dark">Github dimmed (dark)</option>
                      {/* <option value="ncurses">Ncurses (light)</option> */}
                      <option value="coolgray-dark">Cool Gray (dark)</option>
                      <option value="pink-light">Strawberry Milkshake (light)</option>
                    </Form.Control>
                  </Col>
                </Form.Group>
                <VeloForm
                  param={{name: "Downloads Password",
                          description: "Default password to use for downloads",
                          type: "string"}}
                  value={this.state.default_password}
                  setValue={value=>this.setState({default_password: value})}
                />
              </Modal.Body>
              <Modal.Footer>
                <Button variant="secondary" onClick={()=>{
                    this.saveSettings();
                    this.props.onClose();
                }}>
                  Save
                </Button>
              </Modal.Footer>
            </Modal>
        );
    };
}


export default class UserLabel extends React.Component {
    static contextType = UserConfig;
    static propTypes = {

    };

    state = {
        showUserSettings: false,
    }

    componentDidMount() {
        this.source = axios.CancelToken.source();
    }

    componentWillUnmount() {
        this.source.cancel("unmounted");
    }

    setSettings = (options) => {
        // Set the ACE theme according to the theme so they match.
        let ace_options = JSON.parse(this.context.traits.ui_settings || "{}");
        ace_options.fontSize = "18px";
        if (options.theme === "no-theme") {
            ace_options.theme = "ace/theme/xcode";
        } else if (options.theme === "veloci-dark") {
            ace_options.theme = "ace/theme/terminal";
        } else if(options.theme === "veloci-light") {
            ace_options.theme = "ace/theme/xcode";
        } else if(options.theme === "pink-light") {
            ace_options.theme = "ace/theme/xcode";
        } else if(options.theme === "ncurses") {
            ace_options.theme = "ace/theme/github";
        } else if(options.theme === "github-dimmed-dark") {
          ace_options.theme = "ace/theme/vibrant_ink";
        } else if(options.theme === "github-dimmed-light") {
          ace_options.theme = "ace/theme/sqlserver";
        } else if(options.theme === "coolgray-dark") {
          ace_options.theme = "ace/theme/nord_dark";
        }
        options.options = JSON.stringify(ace_options);

        api.post("v1/SetGUIOptions", options, this.source.token).then((response) => {
            this.context.updateTraits();
        });
    }

    render() {
        return (
            <>
              { this.state.showUserSettings &&
                <UserSettings
                  setSetting={this.setSettings}
                  onClose={()=>this.setState({showUserSettings: false})} /> }
              <ButtonGroup className="user-label">
                <Button href={api.base_path + "/app/logoff.html?username="+
                              this.context.traits.username } >
                  <FontAwesomeIcon icon="sign-out-alt" />
                </Button>
                <Button variant="default"
                        onClick={()=>this.setState({showUserSettings: true})}
                >
                  { this.context.traits.username }
                  { this.context.traits.picture &&
                    <img className="toolbar-buttons img-thumbnail"
                         alt="avatar"
                         src={ this.context.traits.picture}
                    />
                  }
                </Button>
              </ButtonGroup>
            </>
        );
    }
};
