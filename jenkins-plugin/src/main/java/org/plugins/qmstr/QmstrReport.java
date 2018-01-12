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
import org.plugins.qmstr.QmstrHttpClient.QmstrHttpClientExeption;

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

        QmstrHttpClient qmstr = new QmstrHttpClient("http://localhost:9000");

        JSONObject linkedTargets = qmstr.linkedTargets();
        if (linkedTargets == null) {
            return false;
        }

        JSONArray linkedtargetsArray = linkedTargets.getJSONArray("linkedtargets");

        Map<String, String> map = new HashMap();
        for (int i=0; i< linkedtargetsArray.size(); i++){

            String targetName = linkedtargetsArray.get(i).toString();
            JSONObject reporttargetNameSpecific  = qmstr.report(targetName);

            map.put(targetName, reporttargetNameSpecific.getString("report"));
        }

        build.addAction(new BuildReportAction(map));
        
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
}
