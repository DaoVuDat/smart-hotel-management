# Tiltfile
# Hotel Management System Development Environment
allow_k8s_contexts('k3d-k3s-cluster')

# Load Kubernetes YAML files
k8s_yaml([
    'k8s/development/namespace.yaml',
    'k8s/development/configmap.yaml',
    'k8s/development/postgres.yaml',
    'k8s/development/secrets.yaml',
    # 'k8s/development/mongodb.yaml',
    # 'k8s/development/redis.yaml',
    # 'k8s/development/kafka.yaml',
    # 'k8s/development/keycloak.yaml',
    # 'k8s/development/kong.yaml',
  
])

# User Service (Go)
docker_build(
    'user-service',
    './services/user-service',
    dockerfile='./services/user-service/Dockerfile',
    live_update=[
        sync('./services/user-service', '/app'),
        run('go build -o /app/main /app/cmd/main.go', trigger=['**/*.go'])
    ]
)

k8s_yaml('k8s/development/user-service.yaml')
k8s_resource(
    'user-service', 
    port_forwards='8001:8001', 
    labels='services',
    )

# Booking Service (Python)
# docker_build(
#     'booking-service',
#     './services/booking-service',
#     dockerfile='./services/booking-service/Dockerfile',
#     live_update=[
#         sync('./services/booking-service/app', '/app/app'),
#         run('pip install -r requirements.txt', trigger=['requirements.txt'])
#     ]
# )

# k8s_yaml('services/booking-service/k8s.yaml')
# k8s_resource('booking-service', port_forwards='8002:8080')

# Room Service (Node.js)
# docker_build(
#     'room-service',
#     './services/room-service',
#     dockerfile='./services/room-service/Dockerfile',
#     live_update=[
#         sync('./services/room-service/src', '/app/src'),
#         run('npm install', trigger=['package.json'])
#     ]
# )

# k8s_yaml('services/room-service/k8s.yaml')
# k8s_resource('room-service', port_forwards='8003:8080')

# Payment Service (Python)
# docker_build(
#     'payment-service',
#     './services/payment-service',
#     dockerfile='./services/payment-service/Dockerfile',
#     live_update=[
#         sync('./services/payment-service/app', '/app/app'),
#         run('pip install -r requirements.txt', trigger=['requirements.txt'])
#     ]
# )

# k8s_yaml('services/payment-service/k8s.yaml')
# k8s_resource('payment-service', port_forwards='8004:8080')

# Notification Service (Go)
# docker_build(
#     'notification-service',
#     './services/notification-service',
#     dockerfile='./services/notification-service/Dockerfile',
#     live_update=[
#         sync('./services/notification-service', '/app'),
#         run('go build -o /app/main /app/cmd/main.go', trigger=['**/*.go'])
#     ]
# )

# k8s_yaml('services/notification-service/k8s.yaml')
# k8s_resource('notification-service', port_forwards='8005:8080')

# Migrations
docker_build(
    'user-db-migration',
    './services/user-service/migrations',
    dockerfile='./services/user-service/migrations/Dockerfile',
    live_update=[
        sync('./services/user-service/migrations', '/app/migrations'),
        run(
            'echo "Migration file updated: you may want to re-run goose"',
            trigger=['**/*.sql']
        ),
    ]
)

k8s_yaml('k8s/development/migrations.yaml')
k8s_resource('user-db-migration', resource_deps=['user-db'], labels='migrations')

# Resource Groups for better organization
k8s_resource('user-db', labels='infrastructure', port_forwards=['5432:5432'])
# k8s_resource('mongodb', resource_deps=['namespace'])
# k8s_resource('redis', resource_deps=['namespace'])
# k8s_resource('kafka', resource_deps=['namespace', 'redis'])
# k8s_resource('keycloak', resource_deps=['namespace', 'postgres'])
# k8s_resource('kong', resource_deps=['namespace', 'postgres'])

# Services depend on infrastructure
k8s_resource('user-service', resource_deps=['user-db', 'user-db-migration'])
# k8s_resource('user-service', resource_deps=['postgres', 'keycloak', 'kong'])
# k8s_resource('booking-service', resource_deps=['postgres', 'kafka', 'kong'])
# k8s_resource('room-service', resource_deps=['mongodb', 'kafka', 'kong'])
# k8s_resource('payment-service', resource_deps=['postgres', 'kafka', 'kong'])
# k8s_resource('notification-service', resource_deps=['postgres', 'kafka', 'kong'])

# Port forwards for external access
# k8s_resource('kong', port_forwards='8080:8000')
# k8s_resource('keycloak', port_forwards='8081:8080')
# k8s_resource('postgres', port_forwards='5432:5432')
# k8s_resource('mongodb', port_forwards='27017:27017')
# k8s_resource('redis', port_forwards='6379:6379')
# k8s_resource('kafka', port_forwards='9092:9092')

# Load balancer for production-like setup
# load_balancer_setup = """
# # Install Kong Gateway
# local_resource('install-kong',
#     cmd='helm repo add kong https://charts.konghq.com && helm repo update',
#     deps=['Tiltfile']
# )

# # Wait for services to be ready
# local_resource('wait-for-services',
#     cmd='kubectl wait --for=condition=available --timeout=300s deployment/user-service deployment/booking-service deployment/room-service deployment/payment-service deployment/notification-service',
#     resource_deps=['user-service', 'booking-service', 'room-service', 'payment-service', 'notification-service']
# )
# """

# Development helpers
# local_resource('proto-gen',
#     cmd='make proto-gen',
#     deps=['shared/proto/*.proto'],
#     resource_deps=[])

# local_resource('db-migrate', 
#     cmd='make db-migrate',
#     resource_deps=['postgres', 'mongodb'])

# Hot reload for development
# local_resource('watch-proto',
#     serve_cmd='make watch-proto',
#     deps=['shared/proto/*.proto'])

print("Hotel Management System - Development Environment")
print("=================================================")
# print("Kong Gateway: http://localhost:8080")
# print("Keycloak: http://localhost:8081")
print("User Service: http://localhost:8001")
# print("Booking Service: http://localhost:8002") 
# print("Room Service: http://localhost:8003")
# print("Payment Service: http://localhost:8004")
# print("Notification Service: http://localhost:8005")
print("=================================================")