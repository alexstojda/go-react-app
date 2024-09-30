import {Configuration, DefaultApi} from "./generated";

export class Api {
    private configuration(): Configuration {
        const openapiConfig = new Configuration();
        if (import.meta.env.VITE_API_HOST !== "")
            openapiConfig.basePath = import.meta.env.VITE_API_HOST
        return openapiConfig;
    };

    public api(): DefaultApi {
        return new DefaultApi(this.configuration());
    };
}
