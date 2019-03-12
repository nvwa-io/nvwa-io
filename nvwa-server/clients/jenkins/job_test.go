package jenkins

import (
    "log"
    "testing"
)

const (
    test_domain = "http://192.168.34.10:8080"
)

var xmlConfig = `<project>
    <actions/>
    <description>${{description}}</description>
    <keepDependencies>false</keepDependencies>
    <properties>
        <jenkins.model.BuildDiscarderProperty>
            <strategy class="hudson.tasks.LogRotator">
                <daysToKeep>-1</daysToKeep>
                <numToKeep>7</numToKeep>
                <artifactDaysToKeep>-1</artifactDaysToKeep>
                <artifactNumToKeep>-1</artifactNumToKeep>
            </strategy>
        </jenkins.model.BuildDiscarderProperty>
        <hudson.model.ParametersDefinitionProperty>
            <parameterDefinitions>
                <net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition plugin="git-parameter@0.9.3">
                    <name>GITLAB_BRANCH</name>
                    <description/>
                    <uuid>073cbea8-233b-4fc8-8b26-f9f79d7ebdf0</uuid>
                    <type>PT_BRANCH</type>
                    <branch/>
                    <tagFilter>*</tagFilter>
                    <branchFilter>.*</branchFilter>
                    <sortMode>NONE</sortMode>
                    <defaultValue>master</defaultValue>
                    <selectedValue>NONE</selectedValue>
                    <quickFilterEnabled>false</quickFilterEnabled>
                    <listSize>5</listSize>
                </net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition>
                <hudson.model.StringParameterDefinition>
                    <name>BUILD_JOB_ID</name>
                    <description/>
                    <defaultValue>-1</defaultValue>
                    <trim>false</trim>
                </hudson.model.StringParameterDefinition>
                <hudson.model.StringParameterDefinition>
                    <name>UID</name>
                    <description/>
                    <defaultValue/>
                    <trim>false</trim>
                </hudson.model.StringParameterDefinition>
            </parameterDefinitions>
        </hudson.model.ParametersDefinitionProperty>
    </properties>
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.9.1">
        <configVersion>2</configVersion>
        <userRemoteConfigs>
            <hudson.plugins.git.UserRemoteConfig>
                <url>
                    ${{gitlab_url}}
                </url>
                <credentialsId>${{gitlab_credentialsId}}</credentialsId>
            </hudson.plugins.git.UserRemoteConfig>
        </userRemoteConfigs>
        <branches>
            <hudson.plugins.git.BranchSpec>
                <name>$GITLAB_BRANCH</name>
            </hudson.plugins.git.BranchSpec>
        </branches>
        <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
        <submoduleCfg class="list"/>
        <extensions/>
    </scm>
    <canRoam>true</canRoam>
    <disabled>false</disabled>
    <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
    <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
    <triggers/>
    <concurrentBuild>false</concurrentBuild>
    <builders>
        <hudson.tasks.Shell>
            <command>
export TIMESTAMP="201901010101"
echo REPO_TAG=$GITLAB_BRANCH.${GIT_COMMIT:0:6}.${BUILD_NUMBER}.$TIMESTAMP > hulk-container-profiles
</command>
</hudson.tasks.Shell>
<EnvInjectBuilder plugin="envinject@2.1.6">
<info>
<propertiesFilePath>hulk-container-profiles</propertiesFilePath>
</info>
</EnvInjectBuilder>
<hudson.tasks.Shell>
<command>
echo ${{docker_repo_url}}
echo ${JOB_NAME}
echo "agent dockerContent create --subProjectName=$JOB_NAME"
echo "agent buildjob create --buildJobId=$BUILD_JOB_ID --uid=$UID --repoTag=$REPO_TAG --location=${{docker_repo_url}} --subProjectName=$JOB_NAME"
</command>
</hudson.tasks.Shell>
<com.cloudbees.dockerpublish.DockerBuilder plugin="docker-build-publish@1.3.2">
<server plugin="docker-commons@1.13">
<uri>tcp://127.0.0.1:2376</uri>
</server>
<registry plugin="docker-commons@1.13">
<url>${{docker_repo_url}}</url>
<credentialsId>${{docker_repo_credentialsId}}</credentialsId>
</registry>
<repoName>${{docker_repo_name}}</repoName>
<noCache>false</noCache>
<forcePull>true</forcePull>
<dockerfilePath>hulk-dockerfile</dockerfilePath>
<skipBuild>false</skipBuild>
<skipDecorate>false</skipDecorate>
<repoTag>$REPO_TAG</repoTag>
<skipPush>false</skipPush>
<createFingerprint>true</createFingerprint>
<skipTagLatest>true</skipTagLatest>
<buildAdditionalArgs/>
<forceTag>false</forceTag>
</com.cloudbees.dockerpublish.DockerBuilder>
</builders>
<publishers>
<org.jenkinsci.plugins.postbuildscript.PostBuildScript plugin="postbuildscript@2.7.0">
<config>
<scriptFiles/>
<groovyScripts/>
<buildSteps>
<org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>
<results>
<string>SUCCESS</string>
</results>
<role>BOTH</role>
<buildSteps>
<hudson.tasks.Shell>
<command>
echo "hulk-container-agent buildJob update --subProjectName=$JOB_NAME --buildJobId=${BUILD_JOB_ID} --status=0"
echo "hulk-container-agent buildJob notify --status=0 repoTag=${REPO_TAG}"
</command>
</hudson.tasks.Shell>
</buildSteps>
</org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>
<org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>
<results>
<string>FAILURE</string>
</results>
<role>BOTH</role>
<buildSteps>
<hudson.tasks.Shell>
<command>
echo "hulk-container-agent buildJob update --subProjectName=$JOB_NAME --buildJobId=${BUILD_JOB_ID} --status=1"
echo "hulk-container-agent buildJob notify --status=1 repoTag=${REPO_TAG}"
</command>
</hudson.tasks.Shell>
</buildSteps>
</org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>
</buildSteps>
<markBuildUnstable>false</markBuildUnstable>
</config>
</org.jenkinsci.plugins.postbuildscript.PostBuildScript>
</publishers>
<buildWrappers/>
</project>`

func Test_Create(t *testing.T) {
    err := C().Config(test_domain, "admin", "admin").Job().Create("demo-02", xmlConfig)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func Test_Update(t *testing.T) {
    err := C().Config(test_domain, "admin", "admin").Job().Update("demo-02", xmlConfig)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func Test_IsExist(t *testing.T) {
    isExist, err := C().Config(test_domain, "admin", "admin").Job().IsExist("demo-02")
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println(isExist)
}

func Test_Delete(t *testing.T) {
    err := C().Config(test_domain, "admin", "admin").Job().Delete("demo-02")
    if err != nil {
        log.Fatal(err.Error())
    }
}

func Test_GetTextLog(t *testing.T) {
    textLog, err := C().Config(test_domain, "admin", "admin").Job().GetTextLog("demo-01", 1)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println(textLog)
}

func Test_GetStatus(t *testing.T) {
    result, isBuilding, err := C().Config(test_domain, "admin", "admin").Job().GetStatus("demo-01", 1)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println(result, isBuilding)
}
