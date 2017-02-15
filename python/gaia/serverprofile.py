# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class ServerProfile(RESTObject):
    """ Represents a ServerProfile in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a ServerProfile instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> serverprofile = ServerProfile(id=u'xxxx-xxx-xxx-xxx', name=u'ServerProfile')
              >>> serverprofile = ServerProfile(data=my_dict)
        """

        super(ServerProfile, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._iptablesmarkvalue = None
        self._ignoreexpression = None
        self._pubookkeepinginterval = None
        self._puheartbeatinterval = None
        self._annotation = None
        self._associatedtags = None
        self._createdat = None
        self._description = None
        self._dockersocketaddress = None
        self._dockersockettype = None
        self._kubernetesenable = None
        self._name = None
        self._namespace = None
        self._normalizedtags = None
        self._parentid = None
        self._parenttype = None
        self._policysynchronizationinterval = None
        self._protected = None
        self._proxylistenaddress = None
        self._receivernumberofqueues = None
        self._receiverqueue = None
        self._receiverqueuesize = None
        self._remoteenforcer = None
        self._status = None
        self._supportlinuxprocesses = None
        self._targetnetworks = None
        self._transmitternumberofqueues = None
        self._transmitterqueue = None
        self._transmitterqueuesize = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="IPTablesMarkValue", remote_name="IPTablesMarkValue")
        self.expose_attribute(local_name="IgnoreExpression", remote_name="IgnoreExpression")
        self.expose_attribute(local_name="PUBookkeepingInterval", remote_name="PUBookkeepingInterval")
        self.expose_attribute(local_name="PUHeartbeatInterval", remote_name="PUHeartbeatInterval")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="description", remote_name="description")
        self.expose_attribute(local_name="dockerSocketAddress", remote_name="dockerSocketAddress")
        self.expose_attribute(local_name="dockerSocketType", remote_name="dockerSocketType")
        self.expose_attribute(local_name="kubernetesEnable", remote_name="kubernetesEnable")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="normalizedTags", remote_name="normalizedTags")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="policySynchronizationInterval", remote_name="policySynchronizationInterval")
        self.expose_attribute(local_name="protected", remote_name="protected")
        self.expose_attribute(local_name="proxyListenAddress", remote_name="proxyListenAddress")
        self.expose_attribute(local_name="receiverNumberOfQueues", remote_name="receiverNumberOfQueues")
        self.expose_attribute(local_name="receiverQueue", remote_name="receiverQueue")
        self.expose_attribute(local_name="receiverQueueSize", remote_name="receiverQueueSize")
        self.expose_attribute(local_name="remoteEnforcer", remote_name="remoteEnforcer")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="supportLinuxProcesses", remote_name="supportLinuxProcesses")
        self.expose_attribute(local_name="targetNetworks", remote_name="targetNetworks")
        self.expose_attribute(local_name="transmitterNumberOfQueues", remote_name="transmitterNumberOfQueues")
        self.expose_attribute(local_name="transmitterQueue", remote_name="transmitterQueue")
        self.expose_attribute(local_name="transmitterQueueSize", remote_name="transmitterQueueSize")
        self.expose_attribute(local_name="updatedAt", remote_name="updatedAt")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        
        return self.ID
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        
        self.ID = ID
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return serverprofileIdentity

    # Properties
    @property
    def ID(self):
        """ Get ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        return self._id

    @ID.setter
    def ID(self, value):
        """ Set ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        self._id = value
    
    @property
    def IPTablesMarkValue(self):
        """ Get IPTablesMarkValue value.

          Notes:
              IptablesMarkValue is the mark value to be used in an iptables implementation.

              
        """
        return self._iptablesmarkvalue

    @IPTablesMarkValue.setter
    def IPTablesMarkValue(self, value):
        """ Set IPTablesMarkValue value.

          Notes:
              IptablesMarkValue is the mark value to be used in an iptables implementation.

              
        """
        self._iptablesmarkvalue = value
    
    @property
    def IgnoreExpression(self):
        """ Get IgnoreExpression value.

          Notes:
              IgnoreExpression allows to set a tag expression that will make Aporeto to ignore docker container started with labels matching the rule.

              
        """
        return self._ignoreexpression

    @IgnoreExpression.setter
    def IgnoreExpression(self, value):
        """ Set IgnoreExpression value.

          Notes:
              IgnoreExpression allows to set a tag expression that will make Aporeto to ignore docker container started with labels matching the rule.

              
        """
        self._ignoreexpression = value
    
    @property
    def PUBookkeepingInterval(self):
        """ Get PUBookkeepingInterval value.

          Notes:
              PUBookkeepingInterval configures how often the PU will be synchronized.

              
        """
        return self._pubookkeepinginterval

    @PUBookkeepingInterval.setter
    def PUBookkeepingInterval(self, value):
        """ Set PUBookkeepingInterval value.

          Notes:
              PUBookkeepingInterval configures how often the PU will be synchronized.

              
        """
        self._pubookkeepinginterval = value
    
    @property
    def PUHeartbeatInterval(self):
        """ Get PUHeartbeatInterval value.

          Notes:
              PUHeartbeatInterval configures the heart beat interval.

              
        """
        return self._puheartbeatinterval

    @PUHeartbeatInterval.setter
    def PUHeartbeatInterval(self, value):
        """ Set PUHeartbeatInterval value.

          Notes:
              PUHeartbeatInterval configures the heart beat interval.

              
        """
        self._puheartbeatinterval = value
    
    @property
    def annotation(self):
        """ Get annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        return self._annotation

    @annotation.setter
    def annotation(self, value):
        """ Set annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        self._annotation = value
    
    @property
    def associatedTags(self):
        """ Get associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        return self._associatedtags

    @associatedTags.setter
    def associatedTags(self, value):
        """ Set associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        self._associatedtags = value
    
    @property
    def createdAt(self):
        """ Get createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        return self._createdat

    @createdAt.setter
    def createdAt(self, value):
        """ Set createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        self._createdat = value
    
    @property
    def description(self):
        """ Get description value.

          Notes:
              Description is the description of the object.

              
        """
        return self._description

    @description.setter
    def description(self, value):
        """ Set description value.

          Notes:
              Description is the description of the object.

              
        """
        self._description = value
    
    @property
    def dockerSocketAddress(self):
        """ Get dockerSocketAddress value.

          Notes:
              DockerSocketAddress is the address of the docker daemon.

              
        """
        return self._dockersocketaddress

    @dockerSocketAddress.setter
    def dockerSocketAddress(self, value):
        """ Set dockerSocketAddress value.

          Notes:
              DockerSocketAddress is the address of the docker daemon.

              
        """
        self._dockersocketaddress = value
    
    @property
    def dockerSocketType(self):
        """ Get dockerSocketType value.

          Notes:
              DockerSocketType is the type of socket to use to talk to the docker daemon.

              
        """
        return self._dockersockettype

    @dockerSocketType.setter
    def dockerSocketType(self, value):
        """ Set dockerSocketType value.

          Notes:
              DockerSocketType is the type of socket to use to talk to the docker daemon.

              
        """
        self._dockersockettype = value
    
    @property
    def kubernetesEnable(self):
        """ Get kubernetesEnable value.

          Notes:
              kubernetesEnable enables  kubernetes mode for the agent.

              
        """
        return self._kubernetesenable

    @kubernetesEnable.setter
    def kubernetesEnable(self, value):
        """ Set kubernetesEnable value.

          Notes:
              kubernetesEnable enables  kubernetes mode for the agent.

              
        """
        self._kubernetesenable = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the entity

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the entity

              
        """
        self._name = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        self._namespace = value
    
    @property
    def normalizedTags(self):
        """ Get normalizedTags value.

          Notes:
              NormalizedTags contains the list of normalized tags of the entities

              
        """
        return self._normalizedtags

    @normalizedTags.setter
    def normalizedTags(self, value):
        """ Set normalizedTags value.

          Notes:
              NormalizedTags contains the list of normalized tags of the entities

              
        """
        self._normalizedtags = value
    
    @property
    def parentID(self):
        """ Get parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        return self._parentid

    @parentID.setter
    def parentID(self, value):
        """ Set parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        self._parentid = value
    
    @property
    def parentType(self):
        """ Get parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        return self._parenttype

    @parentType.setter
    def parentType(self, value):
        """ Set parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        self._parenttype = value
    
    @property
    def policySynchronizationInterval(self):
        """ Get policySynchronizationInterval value.

          Notes:
              PolicySynchronizationInterval configures how often the policy will be resynchronized.

              
        """
        return self._policysynchronizationinterval

    @policySynchronizationInterval.setter
    def policySynchronizationInterval(self, value):
        """ Set policySynchronizationInterval value.

          Notes:
              PolicySynchronizationInterval configures how often the policy will be resynchronized.

              
        """
        self._policysynchronizationinterval = value
    
    @property
    def protected(self):
        """ Get protected value.

          Notes:
              Protected defines if the object is protected.

              
        """
        return self._protected

    @protected.setter
    def protected(self, value):
        """ Set protected value.

          Notes:
              Protected defines if the object is protected.

              
        """
        self._protected = value
    
    @property
    def proxyListenAddress(self):
        """ Get proxyListenAddress value.

          Notes:
              AgentPort is the port the agent should use to listen for API calls 

              
        """
        return self._proxylistenaddress

    @proxyListenAddress.setter
    def proxyListenAddress(self, value):
        """ Set proxyListenAddress value.

          Notes:
              AgentPort is the port the agent should use to listen for API calls 

              
        """
        self._proxylistenaddress = value
    
    @property
    def receiverNumberOfQueues(self):
        """ Get receiverNumberOfQueues value.

          Notes:
              ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network receiver starting at the ReceiverQueue

              
        """
        return self._receivernumberofqueues

    @receiverNumberOfQueues.setter
    def receiverNumberOfQueues(self, value):
        """ Set receiverNumberOfQueues value.

          Notes:
              ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network receiver starting at the ReceiverQueue

              
        """
        self._receivernumberofqueues = value
    
    @property
    def receiverQueue(self):
        """ Get receiverQueue value.

          Notes:
              ReceiverQueue is the base queue number for traffic from the network.

              
        """
        return self._receiverqueue

    @receiverQueue.setter
    def receiverQueue(self, value):
        """ Set receiverQueue value.

          Notes:
              ReceiverQueue is the base queue number for traffic from the network.

              
        """
        self._receiverqueue = value
    
    @property
    def receiverQueueSize(self):
        """ Get receiverQueueSize value.

          Notes:
              ReceiverQueueSize is the queue size of the receiver

              
        """
        return self._receiverqueuesize

    @receiverQueueSize.setter
    def receiverQueueSize(self, value):
        """ Set receiverQueueSize value.

          Notes:
              ReceiverQueueSize is the queue size of the receiver

              
        """
        self._receiverqueuesize = value
    
    @property
    def remoteEnforcer(self):
        """ Get remoteEnforcer value.

          Notes:
              RemoteEnforcer inidicates whether a single enforcer should be used or a distributed enforcer. True means distributed.

              
        """
        return self._remoteenforcer

    @remoteEnforcer.setter
    def remoteEnforcer(self, value):
        """ Set remoteEnforcer value.

          Notes:
              RemoteEnforcer inidicates whether a single enforcer should be used or a distributed enforcer. True means distributed.

              
        """
        self._remoteenforcer = value
    
    @property
    def status(self):
        """ Get status value.

          Notes:
              Status of an entity

              
        """
        return self._status

    @status.setter
    def status(self, value):
        """ Set status value.

          Notes:
              Status of an entity

              
        """
        self._status = value
    
    @property
    def supportLinuxProcesses(self):
        """ Get supportLinuxProcesses value.

          Notes:
              SupportLinuxProcesses configures support for Linux processes.

              
        """
        return self._supportlinuxprocesses

    @supportLinuxProcesses.setter
    def supportLinuxProcesses(self, value):
        """ Set supportLinuxProcesses value.

          Notes:
              SupportLinuxProcesses configures support for Linux processes.

              
        """
        self._supportlinuxprocesses = value
    
    @property
    def targetNetworks(self):
        """ Get targetNetworks value.

          Notes:
              TargetNetworks is the list of networks that authorization should be applied.

              
        """
        return self._targetnetworks

    @targetNetworks.setter
    def targetNetworks(self, value):
        """ Set targetNetworks value.

          Notes:
              TargetNetworks is the list of networks that authorization should be applied.

              
        """
        self._targetnetworks = value
    
    @property
    def transmitterNumberOfQueues(self):
        """ Get transmitterNumberOfQueues value.

          Notes:
              TransmitterNumberOfQueues is the number of queues for application traffic.

              
        """
        return self._transmitternumberofqueues

    @transmitterNumberOfQueues.setter
    def transmitterNumberOfQueues(self, value):
        """ Set transmitterNumberOfQueues value.

          Notes:
              TransmitterNumberOfQueues is the number of queues for application traffic.

              
        """
        self._transmitternumberofqueues = value
    
    @property
    def transmitterQueue(self):
        """ Get transmitterQueue value.

          Notes:
              TransmitterQueue is the queue number for traffic from the applications starting at the transmitterQueue

              
        """
        return self._transmitterqueue

    @transmitterQueue.setter
    def transmitterQueue(self, value):
        """ Set transmitterQueue value.

          Notes:
              TransmitterQueue is the queue number for traffic from the applications starting at the transmitterQueue

              
        """
        self._transmitterqueue = value
    
    @property
    def transmitterQueueSize(self):
        """ Get transmitterQueueSize value.

          Notes:
              TransmitterQueueSize is the size of the queue for application traffic.

              
        """
        return self._transmitterqueuesize

    @transmitterQueueSize.setter
    def transmitterQueueSize(self, value):
        """ Set transmitterQueueSize value.

          Notes:
              TransmitterQueueSize is the size of the queue for application traffic.

              
        """
        self._transmitterqueuesize = value
    
    @property
    def updatedAt(self):
        """ Get updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        return self._updatedat

    @updatedAt.setter
    def updatedAt(self, value):
        """ Set updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        self._updatedat = value
    

    # serverprofileIdentity represents the Identity of the object
serverprofileIdentity = {"name": "serverprofile", "category": "serverprofiles", "constructor": ServerProfile}