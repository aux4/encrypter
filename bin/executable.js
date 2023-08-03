const colors = require("colors");
const { Engine, Output: out } = require("@aux4/engine");
const Secret = require("../lib/Secret");
const Crypto = require("../lib/Crypto");
const { readStream } = require("../lib/Utils");

const { AUX4_SECRET_KEY } = process.env;
const args = process.argv.slice(2);

const aux4 = {
  profiles: [
    {
      name: "main",
      commands: [
        {
          name: "generate-secret",
          execute: async () => {
            try {
              const secret = await Secret.generateSecret();
              out.print(secret);
            } catch (e) {
              out.println(e.message.red);
              process.exit(1);
            }
          },
          help: {
            text: "Generate a new secret key"
          }
        },
        {
          name: "encrypt",
          execute: async (command, action) => {
            const params = command.params;

            const secret = (await params.secret) || AUX4_SECRET_KEY;
            let text = (await params.text) || action.join(" ");
            if (!text || text.length === 0) {
              text = await readStream(process.stdin);
            }

            const crypto = new Crypto(secret);

            try {
              out.println(crypto.encrypt(text));
            } catch (e) {
              out.println(e.message.red);
              process.exit(1);
            }
          },
          help: {
            text: "Encrypt a text using a secret key",
            variables: [
              {
                name: "secret",
                text: "Secret key"
              },
              {
                name: "text",
                text: "Text to encrypt"
              }
            ]
          }
        },
        {
          name: "decrypt",
          execute: async (command, action) => {
            const params = command.params;

            const secret = (await params.secret) || AUX4_SECRET_KEY;
            let token = (await params.token) || action.join(" ");
            if (!token || token.length === 0) {
              token = await readStream(process.stdin);
            }

            const crypto = new Crypto(secret);

            try {
              out.println(crypto.decrypt(token));
            } catch (e) {
              out.println(e.message.red);
              process.exit(1);
            }
          },
          help: {
            text: "Decrypt a token using a secret key",
            variables: [
              {
                name: "secret",
                text: "Secret key"
              },
              {
                name: "token",
                text: "Token to decrypt"
              }
            ]
          }
        }
      ]
    }
  ]
};

(async () => {
  const engine = new Engine({ aux4 });
  await engine.run(args);
})();
