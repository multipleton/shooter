const mocha = require('mocha');
const assert = require('assert');

mocha.describe('Stub test', () => {
  mocha.it('Should pass', () => {
    assert.strictEqual(true, true);
  });
});
