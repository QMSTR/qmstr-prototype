package org.plugins.qmstr;

import hudson.Extension;
import hudson.model.BuildBadgeAction;
import net.sf.json.JSONObject;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.ProtocolException;
import java.net.URL;

@Extension
public class QmstrBadge implements BuildBadgeAction{

    private String status= "";

    public String getStatus() {
        String result = getQmstrStatus().getString("running");
        if (result.equals("ok")){
            return "Qmstr is running";
        }
        return status;
    }

    public JSONObject getQmstrStatus() {
        URL url = null;
        try {
            url = new URL("http://localhost:9000/health");
        } catch (MalformedURLException e) {
            e.printStackTrace();
        }
        HttpURLConnection con = null;
        try {
            con = (HttpURLConnection) url.openConnection();
        } catch (IOException e) {
            e.printStackTrace();
        }
        try {
            con.setRequestMethod("GET");
        } catch (ProtocolException e) {
            e.printStackTrace();
        }
        con.setRequestProperty("Content-Type", "application/json");

        int status = 0;
        try {
            status = con.getResponseCode();
        } catch (IOException e) {
            e.printStackTrace();
        }

        if (status != 200) {
            throw new NullPointerException("Status " + status);
        }

        BufferedReader in = null;
        try {
            in = new BufferedReader(
                    new InputStreamReader(con.getInputStream()));
        } catch (IOException e) {
            e.printStackTrace();
        }
        String inputLine;
        StringBuffer content = new StringBuffer();
        try {
            while ((inputLine = in.readLine()) != null) {
                content.append(inputLine);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        try {
            in.close();
            con.disconnect();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return JSONObject.fromObject(content.toString());
    }

    @Override
    public String getIconFileName() {
        // TODO Auto-generated method stub
        return null;
    }

    @Override
    public String getDisplayName() {
        // TODO Auto-generated method stub
        return "Message";
    }
    @Override
    public String getUrlName() {
        // TODO Auto-generated method stub
        return null;
    }
}
