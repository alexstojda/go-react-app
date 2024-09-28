import {Configuration, DefaultApi} from "./generated";

export class Api {
    private configuration(): Configuration {
        const openapiConfig = new Configuration();
        if (process.env.REACT_APP_API_HOST !== "")
            openapiConfig.basePath = process.env.REACT_APP_API_HOST
        return openapiConfig;
    };

    public api(): DefaultApi {
        return new DefaultApi(this.configuration());
    };
}
