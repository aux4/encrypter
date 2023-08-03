const crypto = require("crypto");

class Secret {
  static async generateSecret() {
    return new Promise((resolve, reject) => {
      crypto.randomBytes(48, (err, buffer) => {
        if (err) {
          reject(new Error(`Error generating key: ${err.message}`));
        }
        resolve(buffer.toString("hex"));
      });
    });
  }
}

module.exports = Secret;
