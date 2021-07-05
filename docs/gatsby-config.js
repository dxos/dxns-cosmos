//
// Copyright 2020 DXOS.org
//

const themeOptions = require("@dxos/docs-theme/theme-options");

module.exports = {
  pathPrefix: "/dxns",
  plugins: [
    {
      resolve: "gatsby-theme-apollo-docs",
      options: {
        ...themeOptions,
        root: __dirname,
        githubRepo: "dxos/dxns",
        description: "DXOS - The Decentralized Operating System",
        subtitle: "DXOS DXNS",
        sidebarCategories: {
          null: ["index"]
        }
      }
    },
    {
      resolve: "gatsby-source-filesystem",
      options: {
        name: "images",
        path: `${__dirname}/src/assets/img`
      }
    },

    // Image processing
    // https://www.gatsbyjs.org/packages/gatsby-plugin-sharp
    // https://www.gatsbyjs.org/packages/gatsby-transformer-sharp
    // https://github.com/lovell/sharp
    "gatsby-plugin-sharp",
    "gatsby-transformer-sharp"
  ]
};
