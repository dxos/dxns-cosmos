//
// Copyright 2020 DXOS.org
//

// Replace apollo-docs page content with our own
// https://github.com/apollographql/gatsby-theme-apollo/blob/master/packages/gatsby-theme-apollo-docs/src/components/page-content.js
import React from "react";
import { PageContent as SharedPageContent } from '@dxos/docs-theme/dist/src/components/PageContent';

export default function PageContent ({ children, ...rest }) {
  return (
    <SharedPageContent {...rest}>
      {children}
    </SharedPageContent>
  );
}
