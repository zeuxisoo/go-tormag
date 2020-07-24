import deepMerge from 'deepmerge';

import configDev from './app';
import configRel from './app.release';

let config = process.env.NODE_ENV !== "development" ? deepMerge(configDev, configRel) : configDev;

export default config
