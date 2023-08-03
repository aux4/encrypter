const Cryptr = require("cryptr");

const DEFAULT_SECRET =
  "4f0c2ac83f5ee523bdd02447127430dcdd1005303a26062ba05b0a8f5d86f2944abd754c1b8e0cef2bbb948e154b1cf3";

class Crypto {
  constructor(secret = DEFAULT_SECRET) {
    this.cryptr = new Cryptr(secret);
  }

  encrypt(text) {
    return this.cryptr.encrypt(text);
  }

  decrypt(token) {
    return this.cryptr.decrypt(token);
  }
}

module.exports = Crypto;
