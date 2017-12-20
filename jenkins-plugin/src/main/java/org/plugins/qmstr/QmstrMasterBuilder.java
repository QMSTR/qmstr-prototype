package org.plugins.qmstr;

import hudson.Extension;
import hudson.Launcher;
import hudson.model.AbstractBuild;
import hudson.model.AbstractProject;
import hudson.model.BuildListener;
import hudson.tasks.BuildStepDescriptor;
import hudson.tasks.Builder;
import org.kohsuke.stapler.DataBoundConstructor;

import java.io.IOException;


public class QmstrMasterBuilder extends Builder {

    @DataBoundConstructor
    public QmstrMasterBuilder(){
    }

    @Extension
    public static class Descriptor extends BuildStepDescriptor<Builder> {

        @Override
        public boolean isApplicable(Class<? extends AbstractProject> jobType) {
            return true;
        }
        @Override
        public String getDisplayName() {
            return "execute Qmstr-master server";
        }
    }

    @Override
    public boolean perform(AbstractBuild<?, ?> build, Launcher launcher, BuildListener listener) throws InterruptedException, IOException {

        QuartermasterProperty prop = build.getProject().getProperty(QuartermasterProperty.class);
        String pathToQMstrMaster;
        Process process;
        if (prop != null){
            pathToQMstrMaster = prop.getPath();
        } else {
            return false;
        }
        process = Runtime.getRuntime().exec(pathToQMstrMaster);
        return true;
    }

}
