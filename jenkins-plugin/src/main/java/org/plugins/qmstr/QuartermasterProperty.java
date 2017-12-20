package org.plugins.qmstr;

import hudson.Extension;
import hudson.model.*;
import org.kohsuke.stapler.DataBoundConstructor;

import net.sf.json.JSONObject;
import org.kohsuke.stapler.StaplerRequest;
import jenkins.model.ParameterizedJobMixIn;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.logging.Logger;

public class QuartermasterProperty extends JobProperty<Job<?, ?>>  {

    Process qmstr_master;
    String path;

    @DataBoundConstructor
    public QuartermasterProperty(String path) {
        this.path= path;
        LOGGER.info(path);
    }

    public String getPath() {
        return path;
    }

    public String getName() {
        return qmstr_master.toString();
    }

    @Extension
    public static final class DescriptorImpl extends JobPropertyDescriptor {
        /**
         * Used to hide property configuration under checkbox,
         * as of not each job is running with Qmstr build environment
         */
        public static final String QMSTRQ_PROJECT_BLOCK_NAME = "quartermasterProject";

        public boolean isApplicable(Class<? extends Job> jobType) {
            return ParameterizedJobMixIn.ParameterizedJob.class.isAssignableFrom(jobType);
        }

        public String getDisplayName() {
            return QMSTRQ_PROJECT_BLOCK_NAME;
        }

        @Override
        public JobProperty<?> newInstance(StaplerRequest req, JSONObject formData) throws FormException {
            QuartermasterProperty tpp = req.bindJSON(
                    QuartermasterProperty.class,
                    formData.getJSONObject(QMSTRQ_PROJECT_BLOCK_NAME)
            );

            if (tpp == null) {
                LOGGER.fine("Couldn't bind JSON");
                return null;
            }

            return tpp;
        }

    }
    private static final Logger LOGGER = Logger.getLogger(QuartermasterProperty.class.getName());
}

