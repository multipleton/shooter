/* eslint-disable no-unused-vars */
/* eslint-disable no-undef */

const ActionsHandler = {
  handlers: {
    login: () => UI.login(),
    create: () => UI.join(),
    connect: () => UI.join(),
    exit: () => UI.exit(),
    start: () => UI.start(),
    logout: () => UI.logout(),
  },
  handleEvent({ target }) {
    const { action } = target.dataset;
    this.handlers[action]();
  },
};
