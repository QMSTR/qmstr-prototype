package org.plugins.qmstr;

import hudson.Extension;
import hudson.Launcher;
import hudson.model.AbstractBuild;
import hudson.model.AbstractProject;
import hudson.model.BuildListener;
import hudson.tasks.BuildStepDescriptor;
import hudson.tasks.BuildStepMonitor;
import hudson.tasks.Publisher;
import hudson.tasks.Recorder;
import net.sf.json.JSONArray;
import net.sf.json.JSONObject;

import org.kohsuke.stapler.DataBoundConstructor;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.ProtocolException;
import java.net.URL;
import java.util.HashMap;
import java.util.Map;

@Extension
public class QmstrReport extends Recorder {

    @DataBoundConstructor
    public QmstrReport(){

    }

    @Override
    public BuildStepMonitor getRequiredMonitorService() {
        return BuildStepMonitor.NONE;
    }

    @Override
    public boolean perform(AbstractBuild<?, ?> build, Launcher launcher, BuildListener listener) throws InterruptedException, IOException {
        JSONObject qmstrlinkedtargets = getQmstrReport("linkedtargets");
        JSONArray linkedtargetsArray = qmstrlinkedtargets.getJSONArray("linkedtargets");
        Map<String, String> map = new HashMap();
        for (int i=0; i< linkedtargetsArray.size(); i++){
            String targetName = linkedtargetsArray.get(i).toString();
            JSONObject reporttargetNameSpecific  = getQmstrReport("report?id=" + targetName);
            map.put(targetName, reporttargetNameSpecific.getString("report"));
        }
        build.addAction(new BuildReportAction(map));
        build.addAction(new QmstrBadge());
        return true;
    }

    @Override
    public BuildStepDescriptor getDescriptor() {
        return (DescriptorImpl) super.getDescriptor();
    }

    @Extension
    public static class DescriptorImpl extends BuildStepDescriptor<Publisher> {

        public DescriptorImpl() {
            super();
        }

        public DescriptorImpl(Class<? extends Publisher> clazz) {
            super(clazz);
            // TODO Auto-generated constructor stub
        }

        @Override
        public boolean isApplicable(Class<? extends AbstractProject> jobType) {
            return true;
        }

        @Override
        public String getDisplayName() {
            return "Generate reuse badge";
        }

    }

    public JSONObject getQmstrReport(String endpoint) {
        URL url = null;
        try {
            url = new URL("http://localhost:9000/" + endpoint);
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
}
