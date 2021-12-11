/* eslint-disable no-unused-vars */
/* eslint-disable no-undef */

const loginSection = document.getElementById('login-section');
const serversSection = document.getElementById('servers-section');
const serverSection = document.getElementById('server-section');
const roundSection = document.getElementById('round-section');
const accountSection = document.getElementById('account-section');
const field = document.getElementById('field');
const log = document.getElementById('log');

const UI = {
  hide(element) {
    // eslint-disable-next-line no-param-reassign
    element.style.display = 'none';
  },
  show(element) {
    // eslint-disable-next-line no-param-reassign
    element.style.display = '';
  },
  hideAll() {
    this.hide(loginSection);
    this.hide(serversSection);
    this.hide(serverSection);
    this.hide(roundSection);
    this.hide(accountSection);
  },
  init() {
    this.hideAll();
    this.show(loginSection);
  },
  login() {
    this.hideAll();
    this.show(serversSection);
    this.show(accountSection);
  },
  join() {
    this.hideAll();
    this.show(serverSection);
    this.show(accountSection);
  },
  exit() {
    this.hideAll();
    this.show(serversSection);
    this.show(accountSection);
  },
  start() {
    this.hideAll();
    this.show(roundSection);
    this.show(accountSection);
  },
  logout() {
    this.init();
  },
};
