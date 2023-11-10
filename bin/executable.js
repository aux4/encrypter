#!/usr/bin/env node

const colors = require("colors");
const {
  Engine,
  Output: out,
  CommandParameters,
  EnvironmentVariableParameterRetriever,
  DefaultParameterRetriever,
  PromptParameterRetriever,
  ConfigParameterRetriever
} = require("@aux4/engine");
const Secret = require("../lib/Secret");
const Crypto = require("../lib/Crypto");
const { readStream } = require("../lib/Utils");

const config = {
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
          execute: async (params, action) => {
            let text = (await params.text) || action.join(" ");
            if (!text || text.length === 0) {
              text = await readStream(process.stdin);
            }

            const secret = (await params.secret) || undefined;
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
                text: "Secret key",
                default: ""
              },
              {
                name: "text",
                text: "Text to encrypt",
                default: ""
              }
            ]
          }
        },
        {
          name: "decrypt",
          execute: async (params, action) => {
            const secret = (await params.secret) || undefined;
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
                text: "Secret key",
                default: ""
              },
              {
                name: "token",
                text: "Token to decrypt",
                default: ""
              }
            ]
          }
        }
      ]
    }
  ]
};

function defaultCommandParameterFactory(config) {
  const commandParametersFactory = CommandParameters.newInstance();
  commandParametersFactory.register(new EnvironmentVariableParameterRetriever());
  commandParametersFactory.register(ConfigParameterRetriever.with(config));
  commandParametersFactory.register(new DefaultParameterRetriever());
  commandParametersFactory.register(new PromptParameterRetriever());
  return commandParametersFactory;
}

(async () => {
  const engine = new Engine({ aux4: config, commandParametersFactory: defaultCommandParameterFactory(config) });

  const args = process.argv.splice(2);

  try {
    await engine.run(args);
  } catch (e) {
    console.error(e.message.red);
    process.exit(1);
  }
})();
