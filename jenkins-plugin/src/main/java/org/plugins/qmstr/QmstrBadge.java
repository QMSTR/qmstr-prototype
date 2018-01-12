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

        QmstrHttpClient qmstr = new QmstrHttpClient("http://localhost:9000");

        JSONObject health = qmstr.health();
        if (health == null) {
            return "";
        }

        String result = health.getString("running");
        if (result.equals("ok")){
            return "Qmstr is running";
        }
        return status;
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
