package org.plugins.qmstr;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.ProtocolException;
import java.net.URL;
import net.sf.json.JSONObject;

/**
 * QmstrHttpClient
 */
public class QmstrHttpClient {

    private String url;

    public QmstrHttpClient(String url) {
        // TODO: trim trailing slashes
        this.url = url;
    }

    public JSONObject health() {
        try {
            return this.getRequest("/health");
        } catch (QmstrHttpClientExeption e) {
            // do something
        }
        // handle better
        return null;
    }

    public JSONObject linkedTargets() {
        try {
            return this.getRequest("/linkedtargets");
        } catch (QmstrHttpClientExeption e) {
            // do something
        }
        // handle better
        return null;
    }

    public JSONObject report(String id) {
        try {
            return this.getRequest("/report?id=" + id);
        } catch (QmstrHttpClientExeption e) {
            // do something
        }
        // handle better
        return null;
    }

    // TODO: Handle errors better and/or log something
    private JSONObject getRequest(String endpoint) throws QmstrHttpClientExeption {

        try {
            URL url = new URL(this.url + endpoint);
            HttpURLConnection con = (HttpURLConnection) url.openConnection();

            con.setRequestMethod("GET");
            con.setRequestProperty("Content-Type", "application/json");

            int status = con.getResponseCode();

            // TODO: handle better. this is too generic
            if (status != 200) {
                throw new QmstrHttpClientExeption("Something went wrong contacting master. Status " + status);
            }

            BufferedReader in = new BufferedReader(new InputStreamReader(con.getInputStream()));
            String inputLine;
            StringBuffer content = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                content.append(inputLine);
            }

            in.close();
            con.disconnect();

            return JSONObject.fromObject(content.toString());

        } catch (MalformedURLException e) {
            // from url
            e.printStackTrace();
        } catch (ProtocolException e) {
            // from setRequestMethod
            e.printStackTrace();
        } catch (IOException e) {
            // from con, status, in, in.readLine(), close, disconnect
            e.printStackTrace();
        }
        // this looks not reachable
        return null;
    }

    /**
     * QmstrHttpClientExeption
     */
    public class QmstrHttpClientExeption extends Exception {

        public QmstrHttpClientExeption() {
            // TODO Auto-generated constructor stub
        }

        public QmstrHttpClientExeption(String message) {
            super(message);
            // TODO Auto-generated constructor stub
        }

        public QmstrHttpClientExeption(Throwable cause) {
            super(cause);
            // TODO Auto-generated constructor stub
        }

        public QmstrHttpClientExeption(String message, Throwable cause) {
            super(message, cause);
            // TODO Auto-generated constructor stub
        }
    }
}